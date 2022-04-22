import React from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import Navbar from "../components/navbar/Navbar";
import Catalog from "../components/catalog/Catalog";
import Footer from "../components/Footer";
import SignIn from "../components/auth/SignIn";

const SignInPage = () => {
    return (
        <div>
            <CssBaseline />
            <Navbar/>
            <SignIn/>
            <Footer />
        </div>
    );
};

export default SignInPage;