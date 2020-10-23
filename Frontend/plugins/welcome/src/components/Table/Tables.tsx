import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import Registerstore from '../Registerstores';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [registerstores, setRegisterstores] = useState<EntRegisterstore[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getRegisterstores = async () => {
     const res = await api.listRegisterstore({ limit: 10, offset: 0 });
     setLoading(false);
     setRegisterstores(res);
   };
   getRegisterstores();
 }, [loading]);
 
 const deleteUsers = async (id: number) => {
   const res = await api.deleteRegisterstore({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">Drug</TableCell>
           <TableCell align="center">Amount</TableCell>
           <TableCell align="center">Store</TableCell>
           <TableCell align="center">User</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {registerstores.map(item => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.edges?.drug?.name}</TableCell>
             <TableCell align="center">{item.amount}</TableCell>
             <TableCell align="center">{item.edges?.store?.name}</TableCell>
             <TableCell align="center">{item.edges?.user?.name}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteUsers(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
               
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
 
