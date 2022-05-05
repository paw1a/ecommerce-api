import './Catalog.css'
import Filter from "./Filter";
import {Button, Grid, LinearProgress, Pagination} from "@mui/material";
import ProductCard from "./Card";
import {useAxios} from "../../api/api";
import Error from "../Error";
import React from "react";

const Catalog = () => {

    let [products, loaded, error] = useAxios("/products/");

    console.log(products);

    return (
        <React.Fragment>
        {!loaded && <LinearProgress color='inherit'/>}
        {error && <Error message='Failed to load products catalog'/>}
            {loaded && !error &&
                <div className='catalog'>
                    <Filter/>

                    <div className='product-catalog' style={{marginTop: '30px'}}>
                        <Grid container columnSpacing={2}>
                            {products.map(product => (
                                <Grid item xs={12} md={4} key={product.id} style={{marginBottom: '20px'}}>
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
                            <Pagination count={10} shape="rounded"/>
                        </div>
                    </div>
                </div>
            }
        </React.Fragment>
    );
};

export default Catalog;