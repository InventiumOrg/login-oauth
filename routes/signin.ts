import express, { Request, Response } from "express";
import { body } from "express-validator";
import { validateRequest, BadRequestError } from "@ticket101/common";
import jwt from "jsonwebtoken";
import { User } from "../models/users";
import { Password } from "../services/password";

const router = express.Router();

router.post("/api/users/signin",
[
    body("email")
        .isEmail()
        .withMessage("Email must be valid"),
    body("password")
        .trim()
        .notEmpty()
        .withMessage("You must supply a password")
]
,validateRequest, async (req: Request, res: Response) => {
    const { email, password } = req.body;
    const existingUser = await User.findOne({ email });

    if (!existingUser) {
        throw new BadRequestError("Invalid credentials")
    }
    
    const passwordsMatch = await Password.compare(existingUser.password, password)

    if (!passwordsMatch){
        throw new BadRequestError("Invalid credentials")
    }

    // Generate Json Web Token

    const userJwt = jwt.sign(
        {
            user_id: existingUser.id,
            email: existingUser.email
        },
        "asdf"
    )
    req.session = {
        jwt: userJwt
    }

    res.status(200).send(existingUser)
});

export { router as signinRouter }