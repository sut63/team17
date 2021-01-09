import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import {
    Container,
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    TextField,
    Button,
    CardMedia,
  } from '@material-ui/core';

  const StudentUI: FC<{}> = () => {
    return ( 
        <TableContainer component={Paper}>
         <Table>
           <TableHead><b>Student Background</b></TableHead>
           <TableBody> 
             <TableRow>
               <TableCell>
                  <Grid item xs={6}>
                  <b>Title</b>
                  <div>
                  <FormControl variant="outlined">
                  <Select name="title" fullWidth>
                      <MenuItem>
                        Mr.
                      </MenuItem>
                  </Select>
                  </FormControl>
                  </div>
                  </Grid>
                  <Grid item xs={6}>
                  <b>School</b>
                  <div>
                  <TextField variant='outlined' type='string'/>
                  </div>
                  </Grid>
                  <Grid item xs={6}>
                    <b>Province</b>
                    <div>
                    <FormControl variant="outlined">
                    <Select name="province" fullWidth>
                      <MenuItem>
                        Chonburi
                      </MenuItem>
                  </Select>
                  </FormControl>
                    </div>
                  </Grid>
                  <Grid item xs={6}>
                    <b>Postal Code</b>
                    <div>
                    <FormControl variant="outlined">
                    <Select name="postal" disabled fullWidth>
                      <MenuItem>
                        20230
                      </MenuItem>
                  </Select>
                  </FormControl>
                    </div>
                  </Grid>
                </TableCell>
                
                <TableCell>
                  <Grid item xs={6}>
                  <b>First Name</b>
                  <div>
                  <TextField variant='outlined' type='string'/>
                  </div>
                  </Grid>
                  <Grid item xs={6}>
                    <b>Degree</b>
                    <div>
                    <FormControl variant="outlined">
                    <Select name="degree" autoWidth>
                      <MenuItem>
                        High School
                      </MenuItem>
                  </Select>
                  </FormControl>
                    </div>
                  </Grid>
                  <Grid item xs={6}>
                    <b>District</b>
                    <div>
                    <FormControl variant="outlined">
                    <Select name="district" autoWidth>
                      <MenuItem>
                        Sriracha
                      </MenuItem>
                  </Select>
                  </FormControl>
                    </div>
                  </Grid>
                  <Grid item xs={6}>
                  <b>Telephone Number</b>
                  <div>
                  <TextField variant='outlined' type='int'/>
                  </div>
                  </Grid>
                </TableCell>

                <TableCell>
                  <Grid item xs={6}>
                  <b>Last Name</b>
                  <div>
                  <TextField variant='outlined' type='string'/>
                  </div>
                  </Grid>
                  <Grid item xs={6}>
                    <b>Recent Address</b>
                    <div>
                    <TextField variant='outlined' type='string'/>
                    </div>
                  </Grid>
                  <Grid item xs={6}>
                    <b>Subdistrict</b>
                    <div>
                    <FormControl variant="outlined">
                    <Select name="zone" autoWidth>
                      <MenuItem>
                        Buang
                      </MenuItem>
                  </Select>
                  </FormControl>
                    </div>
                  </Grid>
                  <Grid item xs={6}>
                  <b>Email</b>
                  <div>
                  <TextField variant='outlined' type='string'/>
                  </div>
                  </Grid>
                </TableCell>
                </TableRow>
           </TableBody>
           </Table>
       </TableContainer>
  );
};export default StudentUI;