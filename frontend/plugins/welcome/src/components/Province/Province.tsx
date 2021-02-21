import React, { FC, useState, useEffect}  from 'react';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Alert from '@material-ui/lab/Alert';
import Swal from 'sweetalert2'; // alert
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
    Snackbar,
    Avatar,
  } from '@material-ui/core';
import{
  Content,
    InfoCard,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';
import  { DefaultApi }  from '../../api/apis';
import { EntRegion, EntContinent, EntCountry} from '../../api/models/';
import { Cookies } from '../../Cookie';

interface ControllersProvince {
  /**
   * 
   * @type {number}
   * @memberof ControllersProvince
   */
  cont?: number;
  /**
   * 
   * @type {number}
   * @memberof ControllersProvince
   */
  coun?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersProvince
   */
  dist?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersProvince
   */
  post?: string;
  /**
   * 
   * @type {string}
   * @memberof ControllersProvince
   */
  prov?: string;
  /**
   * 
   * @type {number}
   * @memberof ControllersProvince
   */
  regi?: number;
  /**
   * 
   * @type {string}
   * @memberof ControllersProvince
   */
  subd?: string;
}


const ProvinceUI: FC<{}> = () => {
  const api = new DefaultApi();
  const [region,Setregion] = useState<EntRegion[]>([]);
  const [country,Setcountry] = useState<EntCountry[]>([]);
  const [continent,Setcontinent] = useState<EntContinent[]>([]);

  const getRegion = async () => {
      const res = await api.listRegion({ limit: 10, offset: 0 });
      Setregion(res);
    };
  const getCount = async () => {
      const res = await api.listCountry({ limit: 10, offset: 0 });
      Setcountry(res);
    };
  const getconti = async () => {
      const res = await api.listContinent({ limit: 10, offset: 0 });
      Setcontinent(res);
    };
  useEffect(() => {
      getCount();
      getRegion();
      getconti();
  }, []);

  

  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const [P, setP] = React.useState< Partial<ControllersProvince>>({});

  const h = (
    event: React.ChangeEvent<{ name?: string; value: any }>,
  ) => {
    const name = event.target.name as keyof typeof P;
    const { value } = event.target;
    setP({ ...P, [name]: value });
    console.log(P);
  };

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/provinces';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(P),
    };
    console.log(P);

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.status);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }
  
 
  //cookie logout
  var cook = new Cookies()
  var cookieName = cook.GetCookie()

  function Clears() {
      cook.ClearCookie()
      window.location.reload(false)
  }

    return ( 
        <Page theme={pageTheme.home}>
        <Header
          title={'Student Management'}
          subtitle='Student Registration Department'>
            <Avatar alt="Remy Sharp"/>
            <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
            <Button variant="text" color="secondary" size="large"
              onClick={Clears} > Logout </Button>
        </Header>
        <Content>
          <ContentHeader title="Student Background"></ContentHeader>
          <TableContainer component={Paper}>
          <Grid container spacing={1}>
                  <Grid item xs={2}>
                    <b>Continent</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="cont" value={P.cont||''}
                    onChange={h}>
                {continent.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.continent}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>Country</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="coun" value={P.coun||''}
                    onChange={h}>
                {country.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.country}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>Region</b>
                    <div>
                <FormControl variant="outlined" fullWidth>
                <Select name="regi" value={P.regi||''}
                    onChange={h}>
                {region.map((item) => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.name}
                    </MenuItem>
                  );
                })}
                </Select>
                </FormControl>
                </div>
                  </Grid>
                </Grid>

                <Grid container spacing={1}>
                  <Grid item xs={2}>
                    <b>Province</b>
                    <div>
               <TextField name='prov' type='string' value={P.prov||''} onChange={h}/>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>District</b>
                    <div>
                <TextField name='dist' type='string' value={P.dist||''} onChange={h}/>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>Subdistrict</b>
                    <div>
                <TextField name='subd' type='string' value={P.subd||''} onChange={h}/>
                </div>
                  </Grid>
                  <Grid item xs={2}>
                    <b>Post Code</b>
                    <div>
                <TextField name='post' type='string' value={P.post||''} onChange={h}/>
                </div>
                  </Grid>
                </Grid>

                <Button variant="contained" color="primary" disableElevation  onClick={() => {
                    save();
                }} > Save </Button>
          </TableContainer>
          </Content>
        </Page>
    );
};export default ProvinceUI;