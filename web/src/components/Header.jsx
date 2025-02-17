import React from "react";
import LanguageSelector from "./LanguageSelector";

export default function Header({ setLanguage, onRunCode }) {
    return (
        <div className="w-full  p-3 flex justify-center items-center gap-4">
            <LanguageSelector setLanguage={setLanguage} />
            <button
                type="button"
                onClick={onRunCode}
                className="flex items-center gap-2 py-2.5 px-5 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 focus:z-10 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
            >
                <span>Run Code</span>
                <svg
                    width="24px"
                    height="24px"
                    viewBox="0 0 16 16"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="#000000"
                >
                    <path d="M2.78 2L2 2.41v12l.78.42 9-6V8l-9-6zM3 13.48V3.35l7.6 5.07L3 13.48z" />
                    <path
                        fillRule="evenodd"
                        clipRule="evenodd"
                        d="M6 14.683l8.78-5.853V8L6 2.147V3.35l7.6 5.07L6 13.48v1.203z"
                    />
                </svg>
            </button>
        </div>
    );
}
