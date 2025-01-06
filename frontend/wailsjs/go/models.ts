export namespace models {
	
	export class Vinculo {
	    id?: number;
	    dni: string;
	    doc_ingreso: number;
	    doc_salida?: number;
	    estado: string;
	    area: number;
	    cargo: number;
	    plaza?: number;
	    regimen: number;
	
	    static createFrom(source: any = {}) {
	        return new Vinculo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.dni = source["dni"];
	        this.doc_ingreso = source["doc_ingreso"];
	        this.doc_salida = source["doc_salida"];
	        this.estado = source["estado"];
	        this.area = source["area"];
	        this.cargo = source["cargo"];
	        this.plaza = source["plaza"];
	        this.regimen = source["regimen"];
	    }
	}
	export class Documento {
	    id?: number;
	    tipoDocumento: string;
	    numeroDocumento?: number;
	    "añoDocumento"?: number;
	    fecha: string;
	    fechaValida?: string;
	    conv?: number;
	    descripcion: string;
	    funcion?: number;
	    sueldo: number;
	
	    static createFrom(source: any = {}) {
	        return new Documento(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.tipoDocumento = source["tipoDocumento"];
	        this.numeroDocumento = source["numeroDocumento"];
	        this["añoDocumento"] = source["añoDocumento"];
	        this.fecha = source["fecha"];
	        this.fechaValida = source["fechaValida"];
	        this.conv = source["conv"];
	        this.descripcion = source["descripcion"];
	        this.funcion = source["funcion"];
	        this.sueldo = source["sueldo"];
	    }
	}
	export class Perfil {
	    Dni: string;
	    Nombre: string;
	    Apaterno: string;
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
	        this.Apaterno = source["Apaterno"];
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
	export class Data {
	    persona: Perfil;
	    documento: Documento;
	    vinculo: Vinculo;
	
	    static createFrom(source: any = {}) {
	        return new Data(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.persona = this.convertValues(source["persona"], Perfil);
	        this.documento = this.convertValues(source["documento"], Documento);
	        this.vinculo = this.convertValues(source["vinculo"], Vinculo);
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

