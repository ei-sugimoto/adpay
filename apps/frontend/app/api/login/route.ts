import { adPayClient } from "@/lib/apiClient";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest): Promise<NextResponse> {
    const body = await req.json();
    const res = await fetch(`${adPayClient.origin}/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
    }).catch((err) => {
        console.error(err);
        return null;
    });

    if (!res) {
        return NextResponse.json({ error: "An error occurred" }, { status: 500 });
    }

    switch (res.status) { 
        case 400:
            return NextResponse.json({ error: "bat request" }, { status: 400 });
        case 401:
            return NextResponse.json({ error: "Invalid credentials" }, { status: 401 });
        case 404:
            return NextResponse.json({ error: "Not found" }, { status: 404 });
        case 500:
            return NextResponse.json({ error: "An error occurred" }, { status: 500 });
        default:
            break;
    }

    const resBody = await res.json();
    const { token } = resBody;
    if (!token) {
        return NextResponse.json({ error: "Invalid credentials" }, { status: 401 });
    }

    const cookieStore = await cookies();
    cookieStore.set("token", token);
    

    return NextResponse.json({ status: 200 });

}