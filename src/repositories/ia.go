package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sashabaranov/go-openai"
)

type IaRepository interface {
	Crear_Query(prompt string) (*string, error)
	Ejecutar_Query(ctx context.Context, query string) ([]map[string]any, error)
	Humanizar_Response(prompt string, respuesta []map[string]any) (*string, error)
}

type iaRepository struct {
	db *sqlx.DB
}

func (i *iaRepository) Humanizar_Response(prompt string, respuesta []map[string]any) (*string, error) {
	client := openai.NewClient(os.Getenv("OPENIA"))

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       openai.GPT4oMini,
		Temperature: 0,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Una persona esta preguntando " + prompt + " y la respuesta de la base de datos es " + fmt.Sprintf("%v", respuesta) + " humaniza esta respuesta y respondela de manera amigable"},
			{Role: "user", Content: prompt},
		},
	})

	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message.Content, nil
}

func (i *iaRepository) Ejecutar_Query(ctx context.Context, query string) ([]map[string]any, error) {

	rows, err := i.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando consulta SQL: %w", err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results []map[string]any
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		rowMap := make(map[string]any)
		for i, colName := range cols {
			rowMap[colName] = columns[i]
		}
		results = append(results, rowMap)
	}
	return results, nil

}

func (i *iaRepository) Crear_Query(prompt string) (*string, error) {
	client := openai.NewClient(os.Getenv("OPENIA"))

	estructuraDB := `
	Estructura de la base de datos:
	CREATE TABLE 'TipoCuenta' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'nombre' UNIQUE ('nombre')
)
ENGINE = InnoDB;
CREATE TABLE 'Plaza' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'codigo' VARCHAR(20) NOT NULL,
  'area_id' INT NOT NULL,
  'cargo_id' INT NOT NULL,
  'estado' ENUM('ocupada','vacante') NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'codigo' UNIQUE ('codigo')
)
ENGINE = InnoDB;
CREATE TABLE 'CuentaBancaria' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'dni_persona' CHAR(8) NOT NULL,
  'numero_cuenta' VARCHAR(20) NOT NULL,
  'tipo_cuenta_id' INT NOT NULL,
  'banco_id' INT NOT NULL,
  'estado' TINYINT NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'numero_cuenta' UNIQUE ('numero_cuenta')
)
ENGINE = InnoDB;
CREATE TABLE 'Persona' ( 
  'dni' CHAR(8) NOT NULL,
  'apaterno' VARCHAR(255) NOT NULL,
  'amaterno' VARCHAR(255) NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
  'direccion' VARBINARY(250) NULL DEFAULT NULL ,
  'telf1' VARBINARY(250) NULL DEFAULT NULL ,
  'telf2' VARBINARY(250) NULL DEFAULT NULL ,
  'email' VARBINARY(250) NULL DEFAULT NULL ,
  'ruc' VARCHAR(11) NULL DEFAULT NULL ,
  'pension_id' INT NULL DEFAULT NULL ,
  'fecha_nacimiento' DATE NOT NULL,
  'sexo' ENUM('F','M') NULL DEFAULT NULL ,
  'foto' LONGTEXT NULL DEFAULT NULL ,
   PRIMARY KEY ('dni')
)
ENGINE = InnoDB;
CREATE TABLE 'vinculo_sindicato' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'vinculo_id' INT NOT NULL,
  'sindicato_id' INT NOT NULL,
  'documento_afiliacion' INT NOT NULL,
  'documento_desafiliacion' INT NULL DEFAULT NULL ,
   PRIMARY KEY ('id'),
  CONSTRAINT 'vinculo_id' UNIQUE ('vinculo_id', 'sindicato_id')
)
ENGINE = InnoDB;
CREATE TABLE 'Pension' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'nombre' UNIQUE ('nombre')
)
ENGINE = InnoDB;
CREATE TABLE 'Banco' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
  'codigo' VARCHAR(10) NOT NULL,
  'estado' TINYINT NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'nombre' UNIQUE ('nombre')
)
ENGINE = InnoDB;
CREATE TABLE 'Cargo' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(100) NOT NULL,
  'activo' TINYINT NULL DEFAULT 1 ,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
CREATE TABLE 'Documento' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'tipo' ENUM('Carta','Resolucion','Acta','Doc.Adm','Contrato','Memorando') NOT NULL,
  'numero' INT NULL DEFAULT NULL ,
  'fecha' DATE NULL DEFAULT NULL ,
  'fecha_valida' DATE NULL DEFAULT NULL ,
  'conv' INT NULL DEFAULT NULL ,
  'descripcion' VARCHAR(255) NULL DEFAULT NULL ,
  'year' INT NULL DEFAULT NULL ,
  'funcion' INT NULL DEFAULT NULL ,
  'sueldo' DOUBLE NULL DEFAULT NULL ,
   PRIMARY KEY ('id'),
  CONSTRAINT 'doc' UNIQUE ('numero', 'year', 'tipo')
)
ENGINE = InnoDB;
CREATE TABLE 'GradoAcademico' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
  'descripcion' TEXT NULL DEFAULT NULL ,
  'nivel' INT NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'nombre' UNIQUE ('nombre')
)
ENGINE = InnoDB;
CREATE TABLE 'asistencia_personal' ( 
  'dni' VARCHAR(8) NOT NULL,
  'hora' TIME NULL DEFAULT NULL ,
  'fecha' DATE NULL DEFAULT NULL 
)
ENGINE = InnoDB;
CREATE TABLE 'Area' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(100) NOT NULL,
  'activo' TINYINT NULL DEFAULT NULL ,
  'nivel' INT NULL DEFAULT NULL ,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
CREATE TABLE 'Regimen' ( 
  'id' INT NOT NULL,
  'nombre' VARCHAR(50) NULL DEFAULT NULL ,
  'estructura' INT NULL DEFAULT NULL ,
  'decreto' VARCHAR(255) NULL DEFAULT NULL ,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
CREATE TABLE 'EventoVinculo' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'vinculo_id' INT NOT NULL,
  'tipo_evento' ENUM('rotacion','suspension') NOT NULL,
  'fecha_inicio' DATE NOT NULL,
  'fecha_fin' DATE NULL DEFAULT NULL ,
  'nueva_area_id' INT NULL DEFAULT NULL ,
  'documento_id' INT NOT NULL,
   PRIMARY KEY ('id'),
  CONSTRAINT 'vinculo_id' UNIQUE ('vinculo_id')
)
ENGINE = InnoDB;
CREATE TABLE 'Vinculo' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'dni' CHAR(8) NOT NULL,
  'doc_ingreso_id' INT NULL DEFAULT NULL ,
  'doc_salida_id' INT NULL DEFAULT NULL ,
  'estado' ENUM('activo','inactivo','pendiente') NOT NULL,
  'area_id' INT NOT NULL,
  'cargo_id' INT NOT NULL,
  'plaza_id' INT NULL DEFAULT NULL ,
  'regimen' INT NULL DEFAULT NULL ,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
CREATE TABLE 'HistorialAcademico' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'dni_persona' CHAR(8) NOT NULL,
  'grado_id' INT NOT NULL,
  'fecha_obtencion' DATE NOT NULL,
  'institucion' VARCHAR(255) NOT NULL,
  'especialidad' VARCHAR(255) NULL DEFAULT NULL ,
  'nro_documento' VARCHAR(50) NULL DEFAULT NULL ,
  'estado' TINYINT NOT NULL,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
CREATE TABLE 'Sindicato' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
CREATE TABLE 'usuario' ( 
  'id' INT AUTO_INCREMENT NOT NULL,
  'nombre' VARCHAR(255) NOT NULL,
  'pass' VARBINARY(500) NOT NULL,
  'email' VARCHAR(255) NOT NULL,
  'nickname' VARCHAR(30) NOT NULL,
  'nivel' INT NOT NULL DEFAULT 1 ,
  'create_at' TIMESTAMP NULL DEFAULT 'current_timestamp()' ,
   PRIMARY KEY ('id')
)
ENGINE = InnoDB;
ALTER TABLE 'Plaza' ADD CONSTRAINT 'Plaza_ibfk_1' FOREIGN KEY ('area_id') REFERENCES 'Area' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Plaza' ADD CONSTRAINT 'Plaza_ibfk_2' FOREIGN KEY ('cargo_id') REFERENCES 'Cargo' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'CuentaBancaria' ADD CONSTRAINT 'CuentaBancaria_ibfk_1' FOREIGN KEY ('dni_persona') REFERENCES 'Persona' ('dni') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'CuentaBancaria' ADD CONSTRAINT 'CuentaBancaria_ibfk_2' FOREIGN KEY ('tipo_cuenta_id') REFERENCES 'TipoCuenta' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'CuentaBancaria' ADD CONSTRAINT 'CuentaBancaria_ibfk_3' FOREIGN KEY ('banco_id') REFERENCES 'Banco' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Persona' ADD CONSTRAINT 'Persona_ibfk_1' FOREIGN KEY ('pension_id') REFERENCES 'Pension' ('id') ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE 'vinculo_sindicato' ADD CONSTRAINT 'fk_vs_2' FOREIGN KEY ('sindicato_id') REFERENCES 'Sindicato' ('id') ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE 'vinculo_sindicato' ADD CONSTRAINT 'fk_vs_3' FOREIGN KEY ('documento_afiliacion') REFERENCES 'Documento' ('id') ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE 'vinculo_sindicato' ADD CONSTRAINT 'fk_vs_1' FOREIGN KEY ('vinculo_id') REFERENCES 'Vinculo' ('id') ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE 'vinculo_sindicato' ADD CONSTRAINT 'fk_vs_4' FOREIGN KEY ('documento_desafiliacion') REFERENCES 'Documento' ('id') ON DELETE RESTRICT ON UPDATE RESTRICT;
ALTER TABLE 'EventoVinculo' ADD CONSTRAINT 'EventoVinculo_ibfk_1' FOREIGN KEY ('vinculo_id') REFERENCES 'Vinculo' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'EventoVinculo' ADD CONSTRAINT 'EventoVinculo_ibfk_2' FOREIGN KEY ('nueva_area_id') REFERENCES 'Area' ('id') ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE 'EventoVinculo' ADD CONSTRAINT 'EventoVinculo_ibfk_3' FOREIGN KEY ('documento_id') REFERENCES 'Documento' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'Vinculo_ibfk_1' FOREIGN KEY ('doc_ingreso_id') REFERENCES 'Documento' ('id') ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'Vinculo_ibfk_4' FOREIGN KEY ('area_id') REFERENCES 'Area' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'fk_regimen' FOREIGN KEY ('regimen') REFERENCES 'Regimen' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'Vinculo_ibfk_2' FOREIGN KEY ('doc_salida_id') REFERENCES 'Documento' ('id') ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'Vinculo_ibfk_5' FOREIGN KEY ('cargo_id') REFERENCES 'Cargo' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'Vinculo_ibfk_3' FOREIGN KEY ('dni') REFERENCES 'Persona' ('dni') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'Vinculo' ADD CONSTRAINT 'Vinculo_ibfk_6' FOREIGN KEY ('plaza_id') REFERENCES 'Plaza' ('id') ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE 'HistorialAcademico' ADD CONSTRAINT 'HistorialAcademico_ibfk_1' FOREIGN KEY ('dni_persona') REFERENCES 'Persona' ('dni') ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE 'HistorialAcademico' ADD CONSTRAINT 'HistorialAcademico_ibfk_2' FOREIGN KEY ('grado_id') REFERENCES 'GradoAcademico' ('id') ON DELETE CASCADE ON UPDATE CASCADE;
  `

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       openai.GPT4oMini,
		Temperature: 0,
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "Eres un asistente que genera exclusivamente consultas SQL basadas en la siguiente estructura de base de datos:\n\n" + estructuraDB + "\n\nTu respuesta debe ser única y exclusivamente la consulta SQL, sin explicaciones, comentarios, formato adicional o texto introductorio."},
			{Role: "user", Content: prompt},
		},
	})

	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message.Content, nil
}

func CreateRepoIa(db *sqlx.DB) IaRepository {
	return &iaRepository{db}
}
