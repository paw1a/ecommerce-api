import React from 'react';
import {Button, Grid} from "@mui/material";

import './Cart.css'
import CartCard from "./CartCard";
import Link from "@mui/material/Link";

const Cart = () => {
    return (
        <div className='cart-container'>
            <Link href="/catalog" color='inherit' underline="none">
                {'Return to catalog'}
            </Link>
            <Grid container columnSpacing={2}>
                <Grid item xs={12} md={8}>
                    <div className='card-header'>
                        <h2>Shopping Cart</h2>
                        <h2>3 items</h2>
                    </div>
                    <div className='cart-content'>
                        <CartCard />
                        <CartCard />
                        <CartCard />
                    </div>
                </Grid>
                <Grid item xs={12} md={4}>
                    <div className='cart-summary'>
                        <h2>Order Summary</h2>
                        <div className='total-price'>
                            <h3>Total cost</h3>
                            <strong>3456$</strong>
                        </div>
                        <Button style={{
                            width: '100%',
                            borderRadius: 0
                        }} color='inherit' variant="outlined">Checkout</Button>
                    </div>
                </Grid>
            </Grid>
        </div>
    );
};

export default Cart;