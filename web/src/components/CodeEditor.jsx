import React, { useEffect, useState } from "react";
import Editor from "@monaco-editor/react";
import { LANGUAGES } from "../constant/snippet";

export default function CodeEditor({ language }) {
    const [value, setValue] = useState("");
    useEffect(() => {
        if (LANGUAGES[language].default) {
            setValue(LANGUAGES[language].default);
        }
    }, [language]);
    return (
        <div className="w-[50%]">
            <Editor
                height="100vh"
                width="100%"
                theme="vs-dark"
                language={language}
                value={value}
                onChange={(value) => setValue(value)}
                options={{
                    scrollBeyondLastLine: true,
                    minimap: { enabled: true, scale: 2, size: "fit" },
                    scrollbar: {
                        vertical: "auto",
                        horizontal: "auto",
                        alwaysConsumeMouseWheel: true,
                    },
                    fontFamily: "Consolas",
                    fontSize: "20px",
                    fontWeight: "semibold",
                    padding: {
                        top: "6px",
                        bottom: "6px",
                    },
                    lineHeight: "1.5px",
                    cursorBlinking: "blink",
                    formatOnPaste: true,
                    fontLigatures: true,
                    wordWrap: "on",
                    bracketPairColorization: {
                        enabled: true,
                    },
                }}
            />
        </div>
    );
}
