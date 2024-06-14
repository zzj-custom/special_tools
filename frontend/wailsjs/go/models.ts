export namespace login {
	
	export class CodeDTO {
	    to: string;
	
	    static createFrom(source: any = {}) {
	        return new CodeDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.to = source["to"];
	    }
	}

}

export namespace netcase {
	
	export class ParseInfoDTO {
	    name: string;
	    content: string;
	
	    static createFrom(source: any = {}) {
	        return new ParseInfoDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.content = source["content"];
	    }
	}
	export class ProcessDTO {
	    path: string[];
	    outPath: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.outPath = source["outPath"];
	    }
	}

}

export namespace response {
	
	export class Reply {
	    code: number;
	    msg: string;
	    result?: any;
	
	    static createFrom(source: any = {}) {
	        return new Reply(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.msg = source["msg"];
	        this.result = source["result"];
	    }
	}

}

