import { BASE_URI } from "./constant/config";
import { LANGUAGES } from "./constant/snippet";
import { encodeBase64, decodeBase64 } from "./util";

class Judge {
    _token = "";
    _result = {
        stdout: "",
        stderr: "",
        status: "",
        time: "",
        memory: "",
        laguage_id: 0,
    };
    async execute(sourceCode, language) {
        await this._run(encodeBase64(sourceCode), LANGUAGES[language].id);
    }
    result() {
        return this._result;
    }
    async _run(sourceCode, languageID) {
        const body = {
            source_code: sourceCode,
            language_id: languageID,
        };
        try {
            const resp = await fetch(BASE_URI, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(body),
            });
            const respBody = await resp.json();
            if (respBody.success) {
                this._token = respBody.data.token;
                await this._pollForResult();
            }
        } catch (err) {
            console.error(err);
        }
    }
    async _pollForResult() {
        let result;
        while (true) {
            result = await this._get();
            this._result = result.data;
            if (result.success && result.data.status !== "pending") {
                break;
            }
            await new Promise((resolve) => setTimeout(resolve, 2000));
        }
    }
    async _get() {
        try {
            const resp = await fetch(BASE_URI + `/${this._token}`, {
                method: "GET",
            });
            return await resp.json();
        } catch (err) {
            console.error(err);
        }
    }
}

export default Judge;
