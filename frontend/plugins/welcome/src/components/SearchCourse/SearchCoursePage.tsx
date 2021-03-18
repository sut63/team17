import React, { useState, useEffect } from 'react';

import { Header, Page, pageTheme, Content } from '@backstage/core';
import {
  TableContainer,
  Paper,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  makeStyles,
  fade,
  Button,
  Avatar,
  InputBase,
} from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import { EntCourse, DefaultApi } from '../../api';
import SearchIcon from '@material-ui/icons/Search';
import { Cookies } from '../../Cookie';

const useStyles = makeStyles(theme => ({
  table: {
    minWidth: 650,
  },
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
    padding: 64,
  },
  search: {
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    backgroundColor: fade(theme.palette.common.white, 0.15),
    '&:hover': {
      backgroundColor: fade(theme.palette.common.white, 0.25),
    },
    marginLeft: 8,
    marginRight: 8,
    width: '420px',
  },
  searchIcon: {
    padding: theme.spacing(0, 2),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  inputRoot: {
    color: 'inherit',
  },
  inputInput: {
    padding: theme.spacing(1, 1, 1, 0),
    paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
    transition: theme.transitions.create('width'),
    width: '400px',
  },
  row: {
    width: '100%',
    display: 'flex',
    justifyContent: 'flex-end',
    alignItems: 'center',
  },
}));

const SearchCoursePage = () => {
  const classes = useStyles();
  const [courses, setCourses] = useState<EntCourse[]>([]);
  const [searchText, setSearchText] = useState('');
  const [filterCourses, setFilterCourse] = useState<EntCourse[]>([]);
  const [isFound, setIsFound] = useState<boolean>(true);
  const [isSearchComplete, setIsSearchComplete] = useState(false);

  const api = new DefaultApi();

  useEffect(() => {
    getCourses();
  }, []);
  const getCourses = async () => {
    let coursesResponse: EntCourse[] = await api.listCourse({});
    setFilterCourse(coursesResponse);
    setCourses(coursesResponse);
  };

  const onSearchTextChange = (e: any) => {
    setSearchText(e.target.value);
  };

  const checkIsFound = (item: any): boolean => {
    let found =
      item.course?.includes(searchText) ||
      item.edges?.courFacu?.faculty?.includes(searchText) ||
      item.edges?.courDegr?.degree?.includes(searchText) ||
      item.edges?.courInst?.institution?.includes(searchText);
    return found;
  };

  const onSearchCourse = () => {
    let temp = courses.filter(item => checkIsFound(item));
    setIsFound(temp.length > 0);
    setFilterCourse(temp);
    if (temp.length > 0) {
      setIsSearchComplete(true);
    }
  };

  const onClearSearch = () => {
    setFilterCourse(courses);
    setIsFound(true);
    setIsSearchComplete(false);
    setSearchText('');
  };

  //cookie logout
  var cook = new Cookies();
  var cookieName = cook.GetCookie();

  function Clears() {
    cook.ClearCookie();
    window.location.reload(false);
  }

  return (
    <>
      <Page theme={pageTheme.home}>
        <Header
          title="ระบบค้นหาข้อมูลหลักสูตร"
          subtitle="เพื่อค้นหาข้อมูลหลักสูตรต่างๆภายในมหาลัย"
        >
          <Avatar alt="Remy Sharp" />
          <div style={{ marginLeft: 10, marginRight: 20 }}>{cookieName}</div>
          <Button
            variant="text"
            color="secondary"
            size="large"
            onClick={Clears}
          >
            {' '}
            Logout{' '}
          </Button>
        </Header>
        
        <Content>
          <div className={classes.root}>
            {isSearchComplete ? (
              <Alert severity="success">
                ค้นหาสำเร็จ พบข้อมูลจำนวน {filterCourses.length} รายการ
              </Alert>
            ) : null}
            {isFound == false ? (
              <Alert severity="error">ไม่พบข้อมูลหลักสูตรนี้</Alert>
            ) : null}
            <div className={classes.row}>
              <h3>ค้นหาหลักสูตร</h3>
              <div className={classes.search}>
                <div className={classes.searchIcon}>
                  <SearchIcon />
                </div>
                <InputBase
                  onChange={onSearchTextChange}
                  value={searchText}
                  placeholder="ชื่อหลักสูตร / ระดับปริญญา / สำนักวิชา / สาขาวิชา"
                  classes={{
                    root: classes.inputRoot,
                    input: classes.inputInput,
                  }}
                  inputProps={{ 'aria-label': 'search' }}
                />
              </div>
              <Button onClick={() => onSearchCourse()} variant="contained">
                ค้นหา
              </Button>
              <Button onClick={() => onClearSearch()} variant="contained">
                ล้างข้อมูลการค้นหา
              </Button>
            </div>
            <TableContainer component={Paper}>
              <Table className={classes.table} aria-label="simple table">
                <TableHead>
                  <TableRow>
                    <TableCell align="center">Course Id</TableCell>
                    <TableCell>Course No.</TableCell>
                    <TableCell>Course Name</TableCell>
                    <TableCell>Degree</TableCell>
                    <TableCell>Faculty</TableCell>
                    <TableCell>Institution</TableCell>
                    <TableCell>Annotation</TableCell>
                    <TableCell>Credit</TableCell>
                  </TableRow>
                </TableHead>

                <TableBody>
                  {filterCourses.map((item, index) => (
                    <TableRow key={item.id}>
                      <TableCell>{item.courseId}</TableCell>
                      <TableCell align="center">{index + 1}</TableCell>
                      <TableCell>{item.course}</TableCell>
                      <TableCell>{item.edges?.courDegr?.degree}</TableCell>
                      <TableCell>{item.edges?.courFacu?.faculty}</TableCell>
                      <TableCell>{item.edges?.courInst?.institution}</TableCell>
                      <TableCell>{item.annotation}</TableCell>
                      <TableCell>{item.credit}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </TableContainer>
          </div>
        </Content>
      </Page>
    </>
  );
};

export default SearchCoursePage;
