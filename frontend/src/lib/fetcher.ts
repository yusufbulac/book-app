export class ApiError extends Error {
    status: number;
    details?: unknown;

    constructor(message: string, status: number, details?: unknown) {
        super(message);
        this.name = "ApiError";
        this.status = status;
        this.details = details;
    }
}

type JsonBody = unknown;

type ReqInit = Omit<RequestInit, "body"> & {
    json?: JsonBody;
};

export async function request<T>(url: string, init: ReqInit = {}) {
    const { json, ...rest } = init;

    const headers = new Headers(rest.headers);
    let body: BodyInit | undefined;

    if (json !== undefined) {
        headers.set("Content-Type", "application/json");
        body = JSON.stringify(json);
    }

    const method = rest.method ?? (json !== undefined ? "POST" : "GET");

    const res = await fetch(url, {
        cache: rest.cache ?? "no-store",
        ...rest,
        method,
        headers,
        body,
        credentials: rest.credentials ?? "omit",
        mode: rest.mode ?? "cors",
    });

    if (!res.ok) {
        const text = await res.text().catch(() => "");
        throw new Error(`HTTP ${res.status} ${res.statusText} - ${text || url}`);
    }

    const ct = res.headers.get("content-type") || "";
    if (ct.includes("application/json")) {
        return (await res.json()) as T;
    }

    return (await res.text()) as T;
}

export function apiBaseUrl() {
    if (typeof window === "undefined") {
        return process.env.BOOKS_API_BASE_URL
            ?? process.env.NEXT_PUBLIC_API_BASE_URL
            ?? "http://localhost:8080";
    }
    return process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";
}
