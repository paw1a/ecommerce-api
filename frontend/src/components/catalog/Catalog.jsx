import React, {useEffect, useState} from 'react';
import './Catalog.css'
import Filter from "./Filter";
import {Button, Grid, Pagination} from "@mui/material";
import axios from "axios";
import ProductCard from "./Card";
import Link from "@mui/material/Link";

const Catalog = () => {
    let [products, setProducts] = useState([]);
    useEffect(() => {
        axios.get('http://52.29.184.51:8080/api/v1/products/')
            .then((resp) => {
                const products = resp.data.data;
                setProducts(products);
            });
    }, [setProducts]);

    console.log(products)

    return (
        <div className='catalog'>
            <Filter/>
            <div className='product-catalog' style={{marginTop: '30px'}}>
                <Grid container columnSpacing={2}>
                    {products.map(product => (
                        <Grid item xs={6} md={4} key={product.id} style={{marginBottom: '20px'}}>
                            <ProductCard product={product}/>
                        </Grid>
                    ))}
                </Grid>
                <Button sx={{
                    width: '100%',
                    backgroundColor: '#e8e8e8',
                    color: 'black',
                    borderRadius: 0,
                    marginBottom: '30px',
                    height: '50px',
                    ":hover": {
                        backgroundColor: '#dcdcdc'
                    }
                }}>Показать еще</Button>
                <div className='pagination-container'>
                    <Pagination count={10} shape="rounded" />
                </div>
            </div>
        </div>
    );
};

export default Catalog;