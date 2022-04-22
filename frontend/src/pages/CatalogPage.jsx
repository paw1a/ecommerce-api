import React from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import Navbar from "../components/navbar/Navbar";
import Catalog from "../components/catalog/Catalog";
import Footer from "../components/Footer";

const CatalogPage = () => {
    return (
        <React.Fragment>
            <CssBaseline />
            <Navbar/>
            <Catalog/>
            <Footer />
        </React.Fragment>
    );
};

export default CatalogPage;