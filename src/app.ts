import express, {Application} from "express";
import morgan from "morgan";

import "dotenv/config"
import {Signale} from "signale";
import proxy from "express-http-proxy";

const app:Application = express();
const signale = new Signale();

app.use(morgan('dev'));
const PORT = 3000;
const GATEWAY = "prueba";

app.use('/api/v1/orders',proxy('http://localhost:3001'));
app.use('/api/v1/products',proxy('http://localhost:8080'));

app.listen(PORT, () => {
    signale.success(`Servicio ${GATEWAY} corriendo en http://localhost:${PORT}`);
});