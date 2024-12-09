export namespace models {
	
	export class Perfil {
	    Dni: string;
	    Nombre: string;
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

}

