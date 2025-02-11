import { adPayClient } from "@/lib/apiClient";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest): Promise<NextResponse> {
    const body = await req.json();

    const cookieStore = await cookies();
    const token = cookieStore.get("token");
    if (!token) {
        return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
    }
    const res = await fetch(`${adPayClient.origin}/project`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token.value}`,
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
        case 201:
            return NextResponse.json({ success: true }, { status: 201 });
        case 400:
            return NextResponse.json({ error: "Invalid request" }, { status: 400 });
        case 404:
            return NextResponse.json({ error: "Not found" }, { status: 404 });
        case 401:
            return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
        default:
            return NextResponse.json({ error: "Internal server error" }, { status: 500 });
    }

}
