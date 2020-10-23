import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบลงทะเบียนยาเข้าคลัง' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={` ${profile.givenName || ''}`}
           
     ></Header>
     <Content>
       <ContentHeader title="ข้อมูลการบันทึก">
         <Link component={RouterLink} to="/Create">
           <Button variant="contained" color="primary">
             Adds
           </Button>
         </Link>         
         
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;

