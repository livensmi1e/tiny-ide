import { useState } from "react";
import CodeEditor from "./components/CodeEditor";
import Header from "./components/Header";
import Output from "./components/Output";
import Judge from "./judgeApi";

function App() {
    const [language, setLanguage] = useState("c");
    const [sourceCode, setSourceCode] = useState("");
    const [stdout, setStdout] = useState("This is stdout content.");
    const [stderr, setStderr] = useState("This is stderr content.");

    const handleRunCode = async () => {
        setStdout("");
        setStderr("");
        const judge = new Judge();
        await judge.execute(sourceCode, language);
        const result = judge.result();
        console.log(result);
        setStdout(result.stdout);
        setStderr(result.stderr);
    };
    return (
        <div className="h-screen w-full flex flex-col gap-4 bg-primary text-white overflow-hidden">
            <Header
                setLanguage={setLanguage}
                setStdout={setStdout}
                setStderr={setStderr}
                onRunCode={handleRunCode}
            />
            <div className="flex gap-1">
                <CodeEditor
                    language={language}
                    sourceCode={sourceCode}
                    setSourceCode={setSourceCode}
                />
                <Output stdout={stdout} stderr={stderr} />
            </div>
        </div>
    );
}

export default App;
