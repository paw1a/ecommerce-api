import React from 'react';
import {Button, Grid} from "@mui/material";
import image from "./catalog/img.png";

import './Product.css'
import IconButton from "@mui/material/IconButton";
import StarIcon from "@mui/icons-material/Star";
import ModeCommentIcon from "@mui/icons-material/ModeComment";
import BookmarkIcon from "@mui/icons-material/Bookmark";

const Product = () => {
    return (
        <div className='container'>
            <Grid container columnSpacing={2}>
                <Grid item xs={12} md={4}>
                    <img src={image} alt="image" className='image'/>
                </Grid>
                <Grid item xs={12} md={8}>
                    <div className='content'>
                        <h1>Name of the product</h1>
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
                        <Button variant='outlined' color='inherit'>Add to cart</Button>
                        <IconButton className='bookmark' disableRipple='true'>
                            <BookmarkIcon className='product-button'/>
                        </IconButton>
                    </div>
                </Grid>
            </Grid>
        </div>
    );
};

export default Product;