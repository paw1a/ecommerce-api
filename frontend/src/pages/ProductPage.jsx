import React from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import Navbar from "../components/navbar/Navbar";
import Product from "../components/Product";
import Footer from "../components/Footer";

const ProductPage = () => {
    return (
        <React.Fragment>
            <CssBaseline />
            <Navbar />
            <Product />
            <Footer />
        </React.Fragment>
    );
};

export default ProductPage;