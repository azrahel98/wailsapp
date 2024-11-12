export namespace sql {
	
	export class Conn {
	
	
	    static createFrom(source: any = {}) {
	        return new Conn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class DBStats {
	    MaxOpenConnections: number;
	    OpenConnections: number;
	    InUse: number;
	    Idle: number;
	    WaitCount: number;
	    WaitDuration: number;
	    MaxIdleClosed: number;
	    MaxIdleTimeClosed: number;
	    MaxLifetimeClosed: number;
	
	    static createFrom(source: any = {}) {
	        return new DBStats(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.MaxOpenConnections = source["MaxOpenConnections"];
	        this.OpenConnections = source["OpenConnections"];
	        this.InUse = source["InUse"];
	        this.Idle = source["Idle"];
	        this.WaitCount = source["WaitCount"];
	        this.WaitDuration = source["WaitDuration"];
	        this.MaxIdleClosed = source["MaxIdleClosed"];
	        this.MaxIdleTimeClosed = source["MaxIdleTimeClosed"];
	        this.MaxLifetimeClosed = source["MaxLifetimeClosed"];
	    }
	}
	export class Row {
	
	
	    static createFrom(source: any = {}) {
	        return new Row(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Rows {
	
	
	    static createFrom(source: any = {}) {
	        return new Rows(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Stmt {
	
	
	    static createFrom(source: any = {}) {
	        return new Stmt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Tx {
	
	
	    static createFrom(source: any = {}) {
	        return new Tx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class TxOptions {
	    Isolation: number;
	    ReadOnly: boolean;
	
	    static createFrom(source: any = {}) {
	        return new TxOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Isolation = source["Isolation"];
	        this.ReadOnly = source["ReadOnly"];
	    }
	}

}

export namespace sqlx {
	
	export class Conn {
	    // Go type: reflectx
	    Mapper?: any;
	
	    static createFrom(source: any = {}) {
	        return new Conn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mapper = this.convertValues(source["Mapper"], null);
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
	export class DB {
	    // Go type: reflectx
	    Mapper?: any;
	
	    static createFrom(source: any = {}) {
	        return new DB(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mapper = this.convertValues(source["Mapper"], null);
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
	export class Stmt {
	    // Go type: reflectx
	    Mapper?: any;
	
	    static createFrom(source: any = {}) {
	        return new Stmt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mapper = this.convertValues(source["Mapper"], null);
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
	export class NamedStmt {
	    Params: string[];
	    QueryString: string;
	    Stmt?: Stmt;
	
	    static createFrom(source: any = {}) {
	        return new NamedStmt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Params = source["Params"];
	        this.QueryString = source["QueryString"];
	        this.Stmt = this.convertValues(source["Stmt"], Stmt);
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
	export class Row {
	    // Go type: reflectx
	    Mapper?: any;
	
	    static createFrom(source: any = {}) {
	        return new Row(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mapper = this.convertValues(source["Mapper"], null);
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
	export class Rows {
	    // Go type: reflectx
	    Mapper?: any;
	
	    static createFrom(source: any = {}) {
	        return new Rows(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mapper = this.convertValues(source["Mapper"], null);
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
	
	export class Tx {
	    // Go type: reflectx
	    Mapper?: any;
	
	    static createFrom(source: any = {}) {
	        return new Tx(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Mapper = this.convertValues(source["Mapper"], null);
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

