import { adPayClient } from "@/lib/apiClient";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest): Promise<NextResponse> {
    const body = await req.json();
    const res = await fetch(`${adPayClient.origin}/register`, {
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

    console.log(res.status);

    switch (res.status) {
        case 200:
            return NextResponse.json({ success: true });
        case 400:
            return NextResponse.json({ error: "Invalid request" }, { status: 400 });
        case 409:
            return NextResponse.json({ error: "User already exists" }, { status: 409 });
        default:
            return NextResponse.json({ error: "An error occurred" }, { status: 500 });
    }
}