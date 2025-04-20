import React, { createContext, useState, useEffect, useContext } from "react";
import { DefaultApi } from "../api/api";
import { Configuration } from "../api/configuration";
import Cookies from "js-cookie";

interface ApiContextProps {
    api: DefaultApi;
    token: string | null;
    setToken: (token: string) => void;
}

const ApiContext = createContext<ApiContextProps | undefined>(undefined);

export const ApiProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [token, setToken] = useState<string | null>(Cookies.get("authToken") || null);

    const api = new DefaultApi(
        new Configuration({
            accessToken: token || undefined,
        })
    );

    useEffect(() => {
        if (token) {
            Cookies.set("authToken", token, { expires: 7 });
        }
    }, [token]);

    return (
        <ApiContext.Provider value={{ api, token, setToken }}>
            {children}
        </ApiContext.Provider>
    );
};

export const useApi = (): ApiContextProps => {
    const context = useContext(ApiContext);
    if (!context) {
        throw new Error("useApi must be used within an ApiProvider");
    }
    return context;
};