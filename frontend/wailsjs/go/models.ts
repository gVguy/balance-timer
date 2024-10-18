export namespace main {
	
	export class Config {
	    intervals: {[key: string]: number};
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.intervals = source["intervals"];
	    }
	}

}

