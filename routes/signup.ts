import express, { Request, Response } from "express";
import { validateRequest, BadRequestError } from "@ticket101/common";
import { body } from "express-validator";
import { User } from "../models/users";
import jwt from "jsonwebtoken";

const router = express.Router();

router.post("/api/users/signup", [
    body("email")
        .isEmail()
        .withMessage("Email must be valid"),
    body("password")
        .notEmpty()
        .withMessage("Password must be supplied")
    ],
    validateRequest,
    async (req: Request, res: Response) => {
        const { email, password } = req.body;
        const existingUser = await User.findOne( { email });

        if (existingUser) {
            throw new BadRequestError("Email in use");
        }

        const user = User.build({ email, password });
        await user.save();

        // Generate JWT

        const userJwt = jwt.sign({
            user_id: user.id,
            email: user.email
        }, "asdf",
        {
            expiresIn: "1h"
        });
        
    req.session = {
        jwt: userJwt
    }

    res.status(201).send(user);
});

export { router as signupRouter }