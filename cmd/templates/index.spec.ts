import type { TestSpecItem } from "@egomobile/http-server";

export const get: TestSpecItem[] = [{
    "ref": Symbol("test #1"),
    "expectations": {
        "status": 200,
        "headers": {
            "content-type": /(text\/plain)/i
        }
    }
}];

export const create: TestSpecItem[] = [{
    "ref": Symbol("test #1"),
    "expectations": {
        "status": 200,
        "headers": {
            "content-type": /(text\/plain)/i
        }
    }
}];

export const update: TestSpecItem[] = [{
    "ref": Symbol("test #1"),
    "expectations": {
        "status": 200,
        "headers": {
            "content-type": /(text\/plain)/i
        }
    }
}];

export const remove: TestSpecItem[] = [{
    "ref": Symbol("test #1"),
    "expectations": {
        "status": 200,
        "headers": {
            "content-type": /(text\/plain)/i
        }
    }
}];
