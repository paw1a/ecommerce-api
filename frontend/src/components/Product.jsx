import React, {useEffect, useState} from 'react';
import {Button, Grid} from "@mui/material";
import image from "./catalog/img.png";

import './Product.css'
import IconButton from "@mui/material/IconButton";
import StarIcon from "@mui/icons-material/Star";
import ModeCommentIcon from "@mui/icons-material/ModeComment";
import BookmarkIcon from "@mui/icons-material/Bookmark";
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import Comment from "./Comment";
import {useParams} from "react-router";
import axios from "axios";

const Product = () => {

    const { productID } = useParams()

    let [product, setProduct] = useState([]);
    useEffect(() => {
        axios.get('http://52.29.184.51:8080/api/v1/products/' + productID)
            .then((resp) => {
                const product = resp.data.data;
                setProduct(product);
            });
    }, [setProduct]);

    let [comments, setComments] = useState([]);
    useEffect(() => {
        axios.get('http://52.29.184.51:8080/api/v1/products/' + productID + '/reviews')
            .then((resp) => {
                const comments = resp.data.data;
                setComments(comments);
            });
    }, [setComments]);

    return (
        <div className='container'>
            <Grid container columnSpacing={2}>
                <Grid item xs={12} md={5}>
                    <img src={image} alt="image" className='image'/>
                </Grid>
                <Grid item xs={12} md={7}>
                    <div className='content'>
                        <h1 style={{marginTop: 0}}>{product.name}</h1>
                        <p>{product.description}</p>
                        <div>
                            <IconButton disableRipple='true' style={{paddingLeft: 0}}>
                                <StarIcon className='product-button'/>
                                {product.totalRating}
                            </IconButton>

                            <IconButton disableRipple='true'>
                                <ModeCommentIcon className='product-button' fontSize='small'/>
                                  <span style={{paddingLeft: '5px'}}>{comments.length}</span>
                            </IconButton>
                        </div>

                        <div className='price'>{product.price} <span style={{color: "gray"}}>$</span></div>
                        <Button variant="outlined" color='inherit' startIcon={
                            <ShoppingCartIcon fontSize='large' />
                        }>
                            Add to cart
                        </Button>
                        <IconButton className='bookmark' disableRipple='true' size='large'>
                            <BookmarkIcon className='product-button'/>
                        </IconButton>
                    </div>
                </Grid>
            </Grid>
            <div className='comments-container'>
                <h2>Comments</h2>
                {comments.map(comment => (
                    <Comment comment={comment}/>
                ))}
            </div>
        </div>
    );
};

export default Product;