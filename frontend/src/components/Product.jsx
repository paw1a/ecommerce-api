import React from 'react';
import {Button, Grid} from "@mui/material";
import image from "./catalog/img.png";

import './Product.css'
import IconButton from "@mui/material/IconButton";
import StarIcon from "@mui/icons-material/Star";
import ModeCommentIcon from "@mui/icons-material/ModeComment";
import BookmarkIcon from "@mui/icons-material/Bookmark";
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import Comment from "./Comment";
import Footer from "./Footer";

const Product = () => {
    return (
        <div>
            <div className='container'>
                <Grid container columnSpacing={2}>
                    <Grid item xs={12} md={5}>
                        <img src={image} alt="image" className='image'/>
                    </Grid>
                    <Grid item xs={12} md={7}>
                        <div className='content'>
                            <h1 style={{marginTop: 0}}>Name of the product</h1>
                            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Architecto culpa debitis doloremque expedita facere illo incidunt, repellendus sint unde voluptate. Amet, aperiam autem corporis debitis eligendi facere harum laborum libero magnam possimus. A beatae, debitis dignissimos esse iure laborum minus nihil nisi nulla obcaecati odio odit perferendis quis sunt, tenetur?</p>
                            <div>
                                <IconButton disableRipple='true' style={{paddingLeft: 0}}>
                                    <StarIcon className='product-button'/>
                                    4.5
                                </IconButton>

                                <IconButton disableRipple='true'>
                                    <ModeCommentIcon className='product-button' fontSize='small'/>
                                    17
                                </IconButton>
                            </div>

                            <div className='price'>12000 <span style={{color: "gray"}}>$</span></div>
                            <Button variant="outlined" color='inherit' startIcon={
                                <ShoppingCartIcon fontSize='large' />
                            }>
                                Add to cart
                            </Button>
                            <IconButton className='bookmark' disableRipple='true' size='large'>
                                <BookmarkIcon className='product-button'/>
                            </IconButton>

                            <div className='comments-container'>
                                <h2>Comments</h2>
                                <Comment/>
                                <Comment/>
                                <Comment/>
                                <Comment/>
                                <Comment/>
                            </div>

                        </div>
                    </Grid>
                </Grid>
            </div>
            <Footer />
        </div>
    );
};

export default Product;