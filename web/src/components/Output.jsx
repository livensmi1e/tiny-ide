import React from "react";
import Editor from "@monaco-editor/react";

export default function Output() {
    return (
        <div className="w-[100%]">
            <Editor
                height="50vh"
                width="100%"
                theme="vs-dark"
                language="text"
                options={{
                    scrollBeyondLastLine: true,
                    minimap: { enabled: false },
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
