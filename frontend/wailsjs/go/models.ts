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
	    sueldo?: number;
	
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
	export class PersonaActivo {
	    Cantidad: number;
	    Activos: number;
	
	    static createFrom(source: any = {}) {
	        return new PersonaActivo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Cantidad = source["Cantidad"];
	        this.Activos = source["Activos"];
	    }
	}
	export class RegimenesCantidad {
	    Cantidad: number;
	    Nombre: string;
	
	    static createFrom(source: any = {}) {
	        return new RegimenesCantidad(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Cantidad = source["Cantidad"];
	        this.Nombre = source["Nombre"];
	    }
	}
	export class ResumenIndicadores {
	    personalregistrado: PersonaActivo;
	    renunciasmes: PersonaActivo[];
	    regimenescantidad: RegimenesCantidad[];
	    sexo: RegimenesCantidad[];
	    sindicatos: RegimenesCantidad[];
	
	    static createFrom(source: any = {}) {
	        return new ResumenIndicadores(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.personalregistrado = this.convertValues(source["personalregistrado"], PersonaActivo);
	        this.renunciasmes = this.convertValues(source["renunciasmes"], PersonaActivo);
	        this.regimenescantidad = this.convertValues(source["regimenescantidad"], RegimenesCantidad);
	        this.sexo = this.convertValues(source["sexo"], RegimenesCantidad);
	        this.sindicatos = this.convertValues(source["sindicatos"], RegimenesCantidad);
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

export namespace openai {
	
	export class ToolCall {
	    index?: number;
	    id?: string;
	    type: string;
	    function: FunctionCall;
	
	    static createFrom(source: any = {}) {
	        return new ToolCall(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.id = source["id"];
	        this.type = source["type"];
	        this.function = this.convertValues(source["function"], FunctionCall);
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
	export class FunctionCall {
	    name?: string;
	    arguments?: string;
	
	    static createFrom(source: any = {}) {
	        return new FunctionCall(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.arguments = source["arguments"];
	    }
	}
	export class ChatMessageImageURL {
	    url?: string;
	    detail?: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatMessageImageURL(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.detail = source["detail"];
	    }
	}
	export class ChatMessagePart {
	    type?: string;
	    text?: string;
	    image_url?: ChatMessageImageURL;
	
	    static createFrom(source: any = {}) {
	        return new ChatMessagePart(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.type = source["type"];
	        this.text = source["text"];
	        this.image_url = this.convertValues(source["image_url"], ChatMessageImageURL);
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
	export class ChatCompletionMessage {
	    role: string;
	    content?: string;
	    refusal?: string;
	    MultiContent: ChatMessagePart[];
	    name?: string;
	    function_call?: FunctionCall;
	    tool_calls?: ToolCall[];
	    tool_call_id?: string;
	
	    static createFrom(source: any = {}) {
	        return new ChatCompletionMessage(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.role = source["role"];
	        this.content = source["content"];
	        this.refusal = source["refusal"];
	        this.MultiContent = this.convertValues(source["MultiContent"], ChatMessagePart);
	        this.name = source["name"];
	        this.function_call = this.convertValues(source["function_call"], FunctionCall);
	        this.tool_calls = this.convertValues(source["tool_calls"], ToolCall);
	        this.tool_call_id = source["tool_call_id"];
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

