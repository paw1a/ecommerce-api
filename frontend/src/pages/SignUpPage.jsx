import React from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import Navbar from "../components/navbar/Navbar";
import Footer from "../components/Footer";
import SignUp from "../components/auth/SignUp";

const SignUpPage = () => {
    return (
        <div>
            <CssBaseline />
            <Navbar/>
            <SignUp/>
            <Footer />
        </div>
    );
};

export default SignUpPage;