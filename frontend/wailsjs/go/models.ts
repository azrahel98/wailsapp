export namespace models {
	
	export class Item {
	    Codigo: string;
	    Nombre: string;
	    Monto: string;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Codigo = source["Codigo"];
	        this.Nombre = source["Nombre"];
	        this.Monto = source["Monto"];
	    }
	}
	export class Boleta {
	    Ruc: string;
	    Empleador: string;
	    Fecha: string;
	    TipoDocIdentidad: string;
	    NumeroDoc: string;
	    Nombre: string;
	    Situacion: string;
	    Ingreso: string;
	    TipoTrabajador: string;
	    RegPensionario: string;
	    Cussp?: string;
	    DiasLab: string;
	    DiasnoLab: string;
	    DiasSubs: string;
	    Condicion: string;
	    Holaslab: string;
	    TipoSuspe: string;
	    MoitvoSuspe: string;
	    DiasSuspe: string;
	    Ingresos: Item[];
	    Descuentos: Item[];
	    AportesTrabajador: Item[];
	    Neto: string;
	    AporteEmpleador: Item[];
	
	    static createFrom(source: any = {}) {
	        return new Boleta(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Ruc = source["Ruc"];
	        this.Empleador = source["Empleador"];
	        this.Fecha = source["Fecha"];
	        this.TipoDocIdentidad = source["TipoDocIdentidad"];
	        this.NumeroDoc = source["NumeroDoc"];
	        this.Nombre = source["Nombre"];
	        this.Situacion = source["Situacion"];
	        this.Ingreso = source["Ingreso"];
	        this.TipoTrabajador = source["TipoTrabajador"];
	        this.RegPensionario = source["RegPensionario"];
	        this.Cussp = source["Cussp"];
	        this.DiasLab = source["DiasLab"];
	        this.DiasnoLab = source["DiasnoLab"];
	        this.DiasSubs = source["DiasSubs"];
	        this.Condicion = source["Condicion"];
	        this.Holaslab = source["Holaslab"];
	        this.TipoSuspe = source["TipoSuspe"];
	        this.MoitvoSuspe = source["MoitvoSuspe"];
	        this.DiasSuspe = source["DiasSuspe"];
	        this.Ingresos = this.convertValues(source["Ingresos"], Item);
	        this.Descuentos = this.convertValues(source["Descuentos"], Item);
	        this.AportesTrabajador = this.convertValues(source["AportesTrabajador"], Item);
	        this.Neto = source["Neto"];
	        this.AporteEmpleador = this.convertValues(source["AporteEmpleador"], Item);
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

