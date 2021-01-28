import React, { FC } from 'react';
import { Typography, Grid, Link } from '@material-ui/core';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardContent from '@material-ui/core/CardContent';

const HeaderCustom = {
  minHeight: '50px',
};

const useStyles = makeStyles({
  root: {
    maxWidth: 345,
  },
});

export type ProfileProps = {
  name: string; 
  id: string;
  system: string;
};

export function CardTeam({ name, id, system }: ProfileProps) {
  const classes = useStyles();
  return (
    <Grid item xs={12} md={3}>
      <Card className={classes.root}>
        <CardActionArea>
          <CardContent>
            <Typography gutterBottom variant="h5" component="h2">
              {system}
            </Typography>
            <Typography gutterBottom variant="h5" component="h2">
              {id} {name}
            </Typography>
          </CardContent>
        </CardActionArea>
      </Card>
    </Grid>
  );
}

const Member: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`ระบบประวัตินักศึกษา`} subtitle={"G17 Member"}></Header>
      <Content>
        <ContentHeader title="สมาชิกในกลุ่ม"></ContentHeader>
        <Grid container>

          <CardTeam name={"นางสาวจีนิณีย์  คุปตวุฒินันท์"} id={"B5901258"} system={"Province"}></CardTeam>
          <CardTeam name={"นางสาวลักขณา  เสนเพ็ง"} id={"B5902439"} system={"Activity"}></CardTeam>
          <CardTeam name={"นางสาวอัครา  ชัยพงษ์"} id={"B5914296"} system={"Professor"}></CardTeam>
          <CardTeam name={"นางสาวณิยนาท  ทองปลิว"} id={"B5911516"} system={"Course"}></CardTeam>
          <CardTeam name={"นายณรงค์  ศิริมา"} id={"B6114688"} system={"Result"}></CardTeam>
          <CardTeam name={"นายชัยธวัช  มหาแก้ว"} id={"B5917396"} system={"Student"}></CardTeam> 
          
        </Grid>
      </Content>
    </Page>
  );
};

export default Member;