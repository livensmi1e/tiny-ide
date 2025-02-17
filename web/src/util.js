import { Buffer } from "buffer";

function encodeBase64(input) {
    return Buffer.from(input).toString("base64");
}

function decodeBase64(input) {
    return Buffer.from(input, "base64").toString();
}

export { encodeBase64, decodeBase64 };
