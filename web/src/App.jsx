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
    const [executionInfo, setExecutionInfo] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const handleRunCode = async () => {
        setStdout("");
        setStderr("");
        setExecutionInfo(null);
        setIsLoading(true);

        const judge = new Judge();
        await judge.execute(sourceCode, language);
        const result = judge.result();

        setStdout(result.stdout);
        setStderr(result.stderr);
        setIsLoading(false);
        setExecutionInfo({
            status: "Run success",
            time: `${result.time}ms`,
            memory: `${result.memory}KB`,
        });
    };
    return (
        <div className="h-screen w-full flex flex-col gap-4 bg-primary text-white overflow-hidden">
            <Header
                setLanguage={setLanguage}
                setStdout={setStdout}
                setStderr={setStderr}
                onRunCode={handleRunCode}
                isLoading={isLoading}
            />
            <div className="flex gap-1">
                <CodeEditor
                    language={language}
                    sourceCode={sourceCode}
                    setSourceCode={setSourceCode}
                />
                <Output stdout={stdout} stderr={stderr} />
            </div>

            {executionInfo && (
                <div className="fixed bottom-2 right-4 z-100 text-white p-3 rounded-lg shadow-lg font-semibold">
                    {executionInfo.status}, {executionInfo.time},{" "}
                    {executionInfo.memory}
                </div>
            )}
        </div>
    );
}

export default App;
