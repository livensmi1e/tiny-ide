import React, { useEffect, useState } from "react";
import Editor from "@monaco-editor/react";
import { LANGUAGES } from "../constant/snippet";
import { OPTIONS } from "../constant/editorOption";

export default function CodeEditor({ language, sourceCode, setSourceCode }) {
    useEffect(() => {
        if (LANGUAGES[language].default) {
            setSourceCode(LANGUAGES[language].default);
        }
    }, [language]);
    return (
        <div className="w-[50%]">
            <Editor
                height="100vh"
                width="100%"
                theme="vs-dark"
                language={language}
                value={sourceCode}
                onChange={(value) => setSourceCode(value)}
                options={OPTIONS}
            />
        </div>
    );
}
