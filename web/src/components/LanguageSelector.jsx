import React, { useState } from "react";
import { LANGUAGES } from "../constant/snippet";

const LANGUAGES_LIST = Object.keys(LANGUAGES);

export default function LanguageSelector({ setLanguage }) {
    const [isOpen, setIsOpen] = useState(false);

    const handleSetLanguage = (selectedLanguage) => {
        setLanguage(selectedLanguage);
        setIsOpen(false);
    };

    return (
        <div className="relative">
            {/* Toggle Button */}
            <button
                onClick={() => setIsOpen(!isOpen)}
                className="text-white bg-blue-700 hover:bg-blue-800  focus:outline-nonefont-medium rounded-lg text-sm px-5 py-2.5 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                type="button"
            >
                Select language{" "}
                <svg
                    className="w-2.5 h-2.5 ms-3"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 10 6"
                >
                    <path
                        stroke="currentColor"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="2"
                        d="m1 1 4 4 4-4"
                    />
                </svg>
            </button>

            {/* Dropdown Menu */}
            {isOpen && (
                <div className="z-40 absolute left-0 mt-2 w-44 bg-white divide-y divide-gray-100 rounded-lg shadow-sm dark:bg-gray-700">
                    <ul className="py-2 text-sm text-gray-700 dark:text-gray-200 font-semibold">
                        {LANGUAGES_LIST.map((language) => {
                            return (
                                <li key={language}>
                                    <a
                                        href="#"
                                        onClick={() =>
                                            handleSetLanguage(language)
                                        }
                                        className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                                    >
                                        {LANGUAGES[language].name}
                                    </a>
                                </li>
                            );
                        })}
                    </ul>
                </div>
            )}
        </div>
    );
}
