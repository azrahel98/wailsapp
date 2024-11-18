export namespace models {
	
	export class Perfil {
	    Dni: string;
	    Nombre: string;
	    Direccion: sql.NullString;
	    Telf1: sql.NullString;
	    Telf2: sql.NullString;
	    Email: sql.NullString;
	    Ruc: sql.NullString;
	    Nacimiento: string;
	    Sexo: sql.NullString;
	    Foto: sql.NullString;
	
	    static createFrom(source: any = {}) {
	        return new Perfil(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Dni = source["Dni"];
	        this.Nombre = source["Nombre"];
	        this.Direccion = this.convertValues(source["Direccion"], sql.NullString);
	        this.Telf1 = this.convertValues(source["Telf1"], sql.NullString);
	        this.Telf2 = this.convertValues(source["Telf2"], sql.NullString);
	        this.Email = this.convertValues(source["Email"], sql.NullString);
	        this.Ruc = this.convertValues(source["Ruc"], sql.NullString);
	        this.Nacimiento = source["Nacimiento"];
	        this.Sexo = this.convertValues(source["Sexo"], sql.NullString);
	        this.Foto = this.convertValues(source["Foto"], sql.NullString);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace sql {
	
	export class NullString {
	    String: string;
	    Valid: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NullString(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.String = source["String"];
	        this.Valid = source["Valid"];
	    }
	}

}

