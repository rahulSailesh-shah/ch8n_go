import { Hono } from "hono";
import { auth } from "./lib/auth"; // path to your auth file

const app = new Hono();

console.log(process.env.PORT);

app
  .on(["POST", "GET"], "/api/auth/*", (c) => auth.handler(c.req.raw))
  .get("/", (c) => {
    return c.text("Hello Hono!");
  });

export default app;
