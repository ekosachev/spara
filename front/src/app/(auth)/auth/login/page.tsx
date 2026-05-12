'use client';

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Field, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { LoginInput, loginSchema } from "@/lib/validation/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";

export default function LoginPage() {
    const form = useForm<LoginInput>({
        resolver: zodResolver(loginSchema),
        defaultValues: {
            email: "",
            password: ""
        },
        mode: "onTouched"
    })

    return (
        <Card>
            <CardHeader>
                <CardTitle>Login to your account</CardTitle>
                <CardDescription>
                    Or register a new one
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form id="form-login" onSubmit={form.handleSubmit(() => { })}>
                    <FieldGroup>
                        <Controller
                            name="email"
                            control={form.control}
                            render={({ field, fieldState }) => (
                                <Field data-invalid={fieldState.invalid}>
                                    <FieldLabel htmlFor="form-login-email">
                                        Email Address
                                    </FieldLabel>
                                    <Input
                                        {...field}
                                        id="form-login-email"
                                        type="email"
                                        aria-invalid={fieldState.invalid}
                                        placeholder="johndoe@example.com"
                                        autoComplete="email"
                                    />
                                    {fieldState.invalid && (
                                        <FieldError errors={[fieldState.error]} />
                                    )}
                                </Field>
                            )}
                        />
                        <Controller
                            name="password"
                            control={form.control}
                            render={({ field, fieldState }) => (
                                <Field data-invalid={fieldState.invalid}>
                                    <FieldLabel htmlFor="form-login-password">
                                        Password
                                    </FieldLabel>
                                    <Input
                                        {...field}
                                        id="form-login-password"
                                        type="password"
                                        aria-invalid={fieldState.invalid}
                                        autoComplete="current-password"
                                    />
                                    {fieldState.invalid && (
                                        <FieldError errors={[fieldState.error]} />
                                    )}
                                </Field>
                            )}
                        />
                    </FieldGroup>
                </form>
            </CardContent>
            <CardFooter>
                <Field orientation="horizontal">
                    <Button type="button" variant="outline" onClick={() => form.reset()}>
                        Reset
                    </Button>
                    <Button type="submit" form="form-login">
                        Log In
                    </Button>
                </Field>
            </CardFooter>
        </Card>
    )
}
