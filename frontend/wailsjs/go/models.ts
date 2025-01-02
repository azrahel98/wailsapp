export namespace models {
	
	export class Documento {
	    id?: number;
	    tipo: string;
	    numero?: number;
	    "año"?: number;
	    fecha: string;
	    fecha_valida?: string;
	    conv?: number;
	    descripcion: string;
	    funcion?: number;
	
	    static createFrom(source: any = {}) {
	        return new Documento(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.tipo = source["tipo"];
	        this.numero = source["numero"];
	        this["año"] = source["año"];
	        this.fecha = source["fecha"];
	        this.fecha_valida = source["fecha_valida"];
	        this.conv = source["conv"];
	        this.descripcion = source["descripcion"];
	        this.funcion = source["funcion"];
	    }
	}
	export class Perfil {
	    Dni: string;
	    Nombre: string;
	    Aparterno: string;
	    Amaterno: string;
	    Direccion?: string;
	    Telf1?: string;
	    Telf2?: string;
	    Email?: string;
	    Ruc?: string;
	    Nacimiento: string;
	    Sexo?: string;
	    Foto?: string;
	
	    static createFrom(source: any = {}) {
	        return new Perfil(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Dni = source["Dni"];
	        this.Nombre = source["Nombre"];
	        this.Aparterno = source["Aparterno"];
	        this.Amaterno = source["Amaterno"];
	        this.Direccion = source["Direccion"];
	        this.Telf1 = source["Telf1"];
	        this.Telf2 = source["Telf2"];
	        this.Email = source["Email"];
	        this.Ruc = source["Ruc"];
	        this.Nacimiento = source["Nacimiento"];
	        this.Sexo = source["Sexo"];
	        this.Foto = source["Foto"];
	    }
	}
	export class PersonDniRequest {
	    nombres: string;
	    apellidoPaterno: string;
	    apellidoMaterno: string;
	
	    static createFrom(source: any = {}) {
	        return new PersonDniRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nombres = source["nombres"];
	        this.apellidoPaterno = source["apellidoPaterno"];
	        this.apellidoMaterno = source["apellidoMaterno"];
	    }
	}

}

