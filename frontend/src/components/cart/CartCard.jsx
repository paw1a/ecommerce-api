import React, {useState} from 'react';

import image from '../catalog/img.png'
import './CartCard.css'
import {Button, ButtonGroup} from "@mui/material";
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import IconButton from "@mui/material/IconButton";

const CartCard = () => {
    const [counter, setCounter] = useState(0);

    return (
        <div className='cart-card'>
            <div className='product-image-name'>
                <img className='cart-image' src={image} alt="image"/>
                <div>
                    <h3 style={{margin: 0}}>Name of the product</h3>
                    1200$
                </div>
            </div>

            <ButtonGroup variant="text" className='counter' color='inherit'>
                <Button onClick={()=> {
                    if (counter < 10)
                        setCounter(counter+1)
                }}>+</Button>

                <Button>{counter}</Button>

                <Button onClick={() => {
                    if(counter > 0)
                        setCounter(counter - 1)
                }}>-</Button>
            </ButtonGroup>

            <div>
                <strong>Total: 12000$</strong>
            </div>

            <IconButton disableRipple='true'>
                <DeleteForeverIcon />
            </IconButton>
        </div>
    );
};

export default CartCard;