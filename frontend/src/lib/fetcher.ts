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

function isRecord(value: unknown): value is Record<string, unknown> {
    return typeof value === "object" && value !== null;
}

function getStringField(obj: Record<string, unknown>, key: string): string | null {
    const v = obj[key];
    return typeof v === "string" ? v : null;
}

export async function request<T>(
    input: string,
    init?: RequestInit & { json?: JsonBody }
): Promise<T> {
    const { json, headers, ...rest } = init ?? {};
    const finalHeaders: HeadersInit = {
        Accept: "application/json",
        ...(json ? { "Content-Type": "application/json" } : {}),
        ...(headers ?? {}),
    };

    const res = await fetch(input, {
        ...rest,
        headers: finalHeaders,
        body: json ? JSON.stringify(json) : init?.body,
    });

    if (!res.ok) {
        let payload: unknown = null;
        try {
            payload = await res.json();
        } catch {
            // no-op
        }

        let message = `HTTP ${res.status} ${res.statusText || ""}`.trim();
        if (isRecord(payload)) {
            const m = getStringField(payload, "message");
            const e = getStringField(payload, "error");
            if (m) message = m;
            else if (e) message = e;
        }

        throw new ApiError(message, res.status, payload);
    }

    try {
        return (await res.json()) as T;
    } catch {
        return {} as T;
    }
}

export function apiBaseUrl() {
    return (
        (typeof window === "undefined"
            ? process.env.BOOKS_API_BASE_URL
            : process.env.NEXT_PUBLIC_API_BASE_URL) || "http://localhost:8080"
    );
}
