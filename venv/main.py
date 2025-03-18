# coding: utf-8
import os
from dotenv import load_dotenv
from langchain_community.utilities import SQLDatabase
# from langchain.prompts import ChatPromptTemplate  # Cambiado de langchain_core
# from langchain.utilities import SQLDatabase      # Cambiado de langchain_community
# from langchain.schema import StrOutputParser     # Cambiado de langchain_core
# from langchain.chains import LLMChain           # Usando LLMChain en lugar de Runnable
# from langchain.chat_models import ChatOpenAI    # Cambiado de langchain_openai

def main():
    # Cargar variables de entorno
    load_dotenv()
    
    # Configurar variables
    api_key = os.getenv("OPENIA")
    db_host = os.getenv("DBHOST")
    db_user = os.getenv("DBUSER")
    db_pass = os.getenv("DBPASS")
    db_database = os.getenv("DBNAME")
    db_secret = os.getenv("DBKEY")

    os.environ["OPENAI_API_KEY"] = api_key

    try:
        # Configurar la base de datos
        db_uri = f"mysql+mysqlconnector://{db_user}:{db_pass}@{db_host}:3306/{db_database}?charset=utf8mb4"
        db = SQLDatabase.from_uri(db_uri)

        # Obtener schema
        schema = db.get_table_info()
        
        print(schema)

        # # Configurar el modelo de lenguaje
        # llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)

        # # Configurar plantilla SQL
        # sql_template = """
        # Basado en este esquema:
        # {schema}
        
        # Para esta pregunta:
        # {question}
        
        # Genera una consulta SQL. Los campos encriptados en la tabla Persona son:
        # direccion, telf1, telf2, email
        # Usa: CAST(AES_DECRYPT(campo, '{key}') AS CHAR)
        # """.format(key=db_secret)

        # # Crear cadena SQL
        # sql_chain = LLMChain(
        #     llm=llm,
        #     prompt=ChatPromptTemplate.from_template(sql_template),
        #     output_parser=StrOutputParser()
        # )

        # # Verificar argumentos
        # if len(sys.argv) < 2:
        #     print("Uso: consulta.exe 'Pregunta en lenguaje natural'")
        #     sys.exit(1)

        # # Ejecutar consulta
        # question = sys.argv[1]
        # sql_query = sql_chain.run(schema=schema, question=question)
        
        # # Ejecutar SQL y obtener resultados
        # results = db.run(sql_query)

        # # Generar respuesta en lenguaje natural
        # response_template = """
        # Basado en los resultados: {results}
        # Para la pregunta: {question}
        
        # Proporciona una respuesta clara y concisa en lenguaje natural.
        # """

        # response_chain = LLMChain(
        #     llm=llm,
        #     prompt=ChatPromptTemplate.from_template(response_template)
        # )

        # final_response = response_chain.run(results=results, question=question)
        # print(final_response)

    except Exception as e:
        print(f"Error: {str(e)}")

if __name__ == "__main__":
    main()