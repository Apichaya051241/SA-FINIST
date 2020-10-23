import React, { useEffect, FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {  
  ContentHeader,
  Content,
 } from '@backstage/core';
import Grid from '@material-ui/core/Grid';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
// import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Select from '@material-ui/core/Select';
import { InputLabel, MenuItem } from '@material-ui/core';
import FormControl from '@material-ui/core/FormControl';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import { EntDrug } from '../../api/models/EntDrug';
import { EntStore } from '../../api/models/EntStore';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 300,
      maxWidth: 300,
      marginTop: '1.5%',
      justifyContent: 'center',
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    title: {
      flexGrow: 1,
    },
    b: {
      display: 'block',
      marginLeft: 'auto',
      marginRight: 'auto',
      width: 160,
      height: 55
    }
  }),
);

interface registerstore {
  drug: number;
  store: number;
  user: number;
  amount: number;

}

const Registerstore: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [registerstores, setRegisterstore] = React.useState<Partial<registerstore>>({});
  const [drugs, setDrugs] = React.useState<EntDrug[]>([]);
  const [stores, setStores] = React.useState<EntStore[]>([]);
  const [users, setUsers] = React.useState<EntUser[]>([]);
  


  const Toast = Swal.mixin({
   
    position: 'center',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Registerstore;
    const { value } = event.target;
    setRegisterstore({ ...registerstores, [name]: +value });
    console.log(registerstores);
  };

  // clear input form
  function clear() {
    setRegisterstore({}); 
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/registerstores';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(registerstores),
    };
    // clear();
    console.log(registerstores); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console


    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
           clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        }
        else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }


  const getDrugs = async () => {
    const res = await http.listDrug({ limit: 10, offset: 0 });
    // setAlert(false);
    setDrugs(res);
  };
  const getStore = async () => {
    const res = await http.listStore({ limit: 10, offset: 0 });
    // setAlert(false);
    setStores(res);
  };
  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };
  // Lifecycle Hooks
  useEffect(() => {
    getDrugs();
    getStore();
    getUsers();
  }, []);



  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบลงทะเบียนยาเข้าคลัง
            </Typography>
        </Toolbar>
      </AppBar>
      <Content>
        <ContentHeader title="บันทึกลงทะเบียนยาเข้าคลัง"> </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={1}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ชื่อยา</InputLabel>
                <Select
                  name="drug"
                  label="ชื่อยา"
                  value={registerstores.drug || ''}
                  onChange={handleChange}    >

                  {drugs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={1}>
              <FormControl className={classes.formControl}>
                <TextField
                  name="amount"
                  label="จำนวนยา"
                  type="number"
                  variant="outlined"
                  value={registerstores.amount || ''}
                  onChange={handleChange} />
              </FormControl>
            </Grid>


            <Grid item xs={1}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ชื่อคลัง</InputLabel>
                <Select
                  name="store"
                  label="ชื่อคลัง"
                  value={registerstores.store || ''}
                  onChange={handleChange}
                >
                  {stores.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={1}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ชื่อผู้ลงทะเบียน</InputLabel>
                <Select
                  name="user"
                  label="ชื่อผู้ลงทะเบียน"
                  value={registerstores.user || ''}
                  onChange={handleChange}
                >
                  {users.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <div className={classes.formControl} >
              <Button className={classes.b}
                component={RouterLink} to="/Table"
                variant="contained"
                color="primary"
                startIcon={<SaveIcon />}
                onClick={save}>
                ลงทะเบียนยา
             </Button>
            </div>           

          </form>
        </div>
      </Content>
    </div>
  );
};

export default Registerstore;
