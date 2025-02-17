import { useState } from "react";
import CodeEditor from "./components/CodeEditor";
import Header from "./components/Header";
import Output from "./components/Output";
import RunInfo from "./components/RunInfo";

function App() {
    const [language, setLanguage] = useState("cpp");
    console.log(language);
    return (
        <div className="h-screen w-full flex flex-col gap-4 bg-primary text-white overflow-hidden">
            <Header setLanguage={setLanguage} />
            <div className="flex gap-1">
                <CodeEditor language={language} />
                <div className="flex flex-col gap-1 w-[50%]">
                    <Output />
                    <RunInfo></RunInfo>
                </div>
            </div>
        </div>
    );
}

export default App;
