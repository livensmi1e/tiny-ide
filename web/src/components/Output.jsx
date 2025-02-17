import React, { useState } from "react";
import Editor from "@monaco-editor/react";
import { OPTIONS } from "../constant/editorOption";

export default function Output({ stdout, stderr }) {
    const [activeTab, setActiveTab] = useState("stdout");

    const tabStyle =
        "p-2.5 px-4 text-sm bg-primary text-gray-300 hover:text-white hover:bg-[#262626]";
    const activeStyle = "text-white bg-editor bg-[#262626]";

    const getContentForTab = () => {
        switch (activeTab) {
            case "stdout":
                return stdout;
            case "stderr":
                return stderr;
            default:
                return "";
        }
    };

    return (
        <div className="w-[50%]">
            <div className="tabs flex items-center gap-2">
                <button
                    onClick={() => setActiveTab("stdout")}
                    className={`
                        ${tabStyle} 
                        ${activeTab === "stdout" ? activeStyle : ""} 
                    `}
                >
                    stdout
                </button>
                <button
                    onClick={() => setActiveTab("stderr")}
                    className={`
                        ${tabStyle} 
                        ${activeTab === "stderr" ? activeStyle : ""} 
                    `}
                >
                    stderr
                </button>
            </div>
            <Editor
                height="100vh"
                width="100%"
                theme="vs-dark"
                language="text"
                value={getContentForTab()}
                options={OPTIONS}
            />
        </div>
    );
}
