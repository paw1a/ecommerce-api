import React, {useState} from 'react';
import './Filter.css'
import {Checkbox, FormControlLabel, List, Paper, Slider} from "@mui/material";

const Filter = () => {
    const [sliderValue, setSliderValue] = useState([0, 100]);
    const [price, setPrice] = useState([0, 120000])

    const updateRange = (e, data) => {
        const from = Math.floor(120000 * (data[0] / 100));
        const to = Math.floor(120000 * (data[1] / 100));
        setPrice([from, to]);
        setSliderValue(data);
    };

    return (
        <div className='filter'>
            <h4>Фильтр</h4>

            <h5>Цена</h5>
            <div className='price-box'>
                <div style={{width: '60%'}}>
                    <div>От</div>
                    <div>{price[0]}</div>
                </div>
                <div style={{width: '50%'}}>
                    <div>До</div>
                    <div>{price[1]}</div>
                </div>
            </div>
            <Slider
                style={{width: '80%', color: 'gray'}}
                getAriaLabel={() => 'Minimum distance'}
                value={sliderValue}
                onChange={updateRange}
                disableSwap
            />

            <h5>Бренд</h5>
            <Paper style={{maxHeight: 200, overflow: 'auto', background: "none", boxShadow: "none"}}>
                <List>
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Adidas" />
                </List>
            </Paper>

            <h5>Пол</h5>
            <FormControlLabel control={<Checkbox size={"small"} />} label="Мужской" />
            <FormControlLabel control={<Checkbox size={"small"} />} label="Женский" />
            <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />

            <h5>Категория</h5>
            <Paper style={{maxHeight: 200, overflow: 'auto', background: "none", boxShadow: "none"}}>
                <List>
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                    <FormControlLabel control={<Checkbox size={"small"} />} label="Унисекс" />
                </List>
            </Paper>
        </div>
    );
};

export default Filter;