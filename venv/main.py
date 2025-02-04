import os
from dotenv import load_dotenv

load_dotenv()

api_key = os.getenv("OPENIA")
db_host = os.getenv("DBHOST")
db_user = os.getenv("DBUSER")
db_pass = os.getenv("DBPASS")
db_database = os.getenv("DBNAME")
db_charset = os.getenv("DBCHARSET")
os.environ["OPENAI_API_KEY"] = api_key
db_secret = os.getenv("DBKEY")

# New-Item -ItemType SymbolicLink -Path "D:\proyectos\wails\wailsapp\venv\.env" -Target "D:\proyectos\wails\wailsapp\.env"

from langchain_core.prompts import ChatPromptTemplate
from langchain_community.utilities import SQLDatabase
from langchain_core.output_parsers import StrOutputParser
from langchain_core.runnables import RunnablePassthrough
from langchain_openai import ChatOpenAI

template = f"""
Genera una consulta SQL basada en el siguiente esquema de base de datos:  
{{schema}}

Pregunta del usuario:  
{{question}} 

Nota: Algunos campos en la base de datos están encriptados con `AES_ENCRYPT` en la tabla Persona: direccion,telf1,telf2,email.  
Para desencriptarlos, usa `CAST(AES_DECRYPT(campo, '{db_secret}') AS CHAR)`, donde `'{db_secret}'` es la clave de desencriptación.  
Genera solo la consulta SQL sin explicaciones adicionales. La consulta debe ser válida y ejecutable.
"""


prompt = ChatPromptTemplate.from_template(template)

db_uri = f"mysql+mysqlconnector://{db_user}:{db_pass}@{db_host}:3306/{db_database}?charset=utf8mb4&collation=utf8mb4_unicode_ci"
db = SQLDatabase.from_uri(db_uri)

def get_schema(_):
    return db.get_table_info()

llm = ChatOpenAI(model="gpt-3.5-turbo",temperature=0)

sql_chain = (
    RunnablePassthrough.assign(schema=get_schema)
    | prompt
    | llm.bind(stop=["\nSQLResult:"])
    | StrOutputParser()
)


template = """
Basado en la siguiente información:

Esquema de la base de datos:
{schema}

Pregunta original:
{question}

Resultados de la consulta:
{response}

Por favor, proporciona una respuesta clara y concisa en lenguaje natural que responda la pregunta del usuario.
Evita mencionar detalles técnicos o la consulta SQL utilizada.
"""
prompt_response = ChatPromptTemplate.from_template(template)

def run_query(query):
    print(query)
    return db.run(query)

full_chain = (
    RunnablePassthrough.assign(query=sql_chain).assign(
        schema=get_schema,
        response=lambda vars: run_query(vars["query"]),
    )
    | prompt_response
    | llm
)

print(full_chain.invoke({"question":"Cuantos trabajadores activos con el cargo de sereno chofer hay"}))