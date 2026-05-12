'use client';

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Field, FieldError, FieldGroup, FieldLabel } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { useAuth } from "@/hooks/useAuth";
import { LoginInput, loginSchema, RegisterInput, registerSchema } from "@/lib/validation/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";

export default function RegistrationPage() {
    const form = useForm<RegisterInput>({
        resolver: zodResolver(registerSchema),
        defaultValues: {
            username: "",
            email: "",
            password: ""
        },
        mode: "onTouched",
    })

    const { registerMutation } = useAuth()

    const onSubmit = async (values: RegisterInput) => {
        try {
            await registerMutation.mutateAsync(values);
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
        } catch (error: any) {
            const status = error.status;
            const message = error.message;

            if (status >= 500) {
                form.setError("root", { type: "custom", message: "Server responded with an error. Try again later" })
            } else {
                form.setError("root", { type: "custom", message: message })
            }
        }
    }

    return (
        <Card>
            <CardHeader>
                <CardTitle>Create a new account</CardTitle>
                <CardDescription>
                    Or log in using an existing one
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form id="form-login" onSubmit={form.handleSubmit(onSubmit)}>
                    <FieldGroup>
                        <Controller
                            name="username"
                            control={form.control}
                            render={({ field, fieldState }) => (
                                <Field data-invalid={fieldState.invalid}>
                                    <FieldLabel htmlFor="form-login-username">
                                        Username
                                    </FieldLabel>
                                    <Input
                                        {...field}
                                        id="form-login-username"
                                        type="text"
                                        aria-invalid={fieldState.invalid}
                                        placeholder="johndoe"
                                        autoComplete="username"
                                    />
                                    {fieldState.invalid && (
                                        <FieldError errors={[fieldState.error]} />
                                    )}
                                </Field>
                            )}
                        />
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
                    <Button type="submit" form="form-login" disabled={registerMutation.isPending}>
                        {registerMutation.isPending ? "Creating the account..." : "Register"}
                    </Button>
                </Field>
            </CardFooter>
        </Card>
    )
}
