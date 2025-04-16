"use client";

import { Formik, Form, Field, ErrorMessage } from "formik";
import * as Yup from "yup";
import styles from "./page.module.css";

export default function Home() {
  return (
    <div
      className={`${styles.page} d-flex justify-content-center align-items-center vh-100 bg-light`}
    >
      <main className="p-4 rounded shadow bg-white" style={{ width: 400 }}>
        <h3 className="mb-4 text-center">Login</h3>
        <Formik
          initialValues={{ username: "", password: "" }}
          validationSchema={Yup.object({
            username: Yup.string().required("Username is required"),
            password: Yup.string().required("Password is required"),
          })}
          onSubmit={(values) => {
            fetch("http://localhost:3000/login", {
              method: "POST",
              headers: { "Content-Type": "application/json" },
              body: JSON.stringify(values),
            })
              .then((res) => {
                if (!res.ok) throw new Error("Login failed");
                return res.json();
              })
              .then((data) => {
                alert("Login success");
                router.push("/dashboard");
              })
              .catch((err) => {
                alert("Login error");
                console.error(err);
              });
          }}
        >
          {() => (
            <Form>
              <div className="mb-3">
                <label htmlFor="username" className="form-label">
                  Username
                </label>
                <Field name="username" type="text" className="form-control" />
                <div className="text-danger small">
                  <ErrorMessage name="username" />
                </div>
              </div>

              <div className="mb-3">
                <label htmlFor="password" className="form-label">
                  Password
                </label>
                <Field
                  name="password"
                  type="password"
                  className="form-control"
                />
                <div className="text-danger small">
                  <ErrorMessage name="password" />
                </div>
              </div>

              <button type="submit" className="btn btn-primary w-100">
                Login
              </button>
            </Form>
          )}
        </Formik>
      </main>
    </div>
  );
}
