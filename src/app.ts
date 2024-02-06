import express from "express";
import { json } from "body-parser";
import { signinRouter } from "../routes/signin";
import { signoutRouter } from "../routes/signout";
import { signupRouter } from "../routes/signup";
import { currentuserRouter } from "../routes/currentuser";
import cookieSession from "cookie-session";
import { NotFoundError, errorHandler } from "@ticket101/common";
import "express-async-errors";

const app = express();

app.set("trust proxy", true);
app.use(json());
app.use(cookieSession({
    signed: false,
    secure: false
}));

app.use(currentuserRouter);
app.use(signinRouter);
app.use(signoutRouter);
app.use(signupRouter);


app.all("*", async (req, res) => {
    throw new NotFoundError();
})

app.use(errorHandler)

export { app }