import { betterAuth } from "better-auth";
import { drizzleAdapter } from "better-auth/adapters/drizzle";
import { openAPI, jwt } from "better-auth/plugins";
import { db } from "@/db/db";

export const auth = betterAuth({
  database: drizzleAdapter(db, {
    provider: "pg",
  }),
  trustedOrigins: ["http://localhost:5173", "http://127.0.0.1:5173"],
  emailAndPassword: {
    enabled: true,
    autoSignIn: true,
  },
  plugins: [openAPI(), jwt()],
});
