// import styled from 'styled-components';

// export const Test = styled.div`
//  display: flex;
// `;
//
import { makeStyles } from '@material-ui/core/styles';

export const useStyles = makeStyles(theme => ({
    root: {
      width: '100%',
      marginTop: theme.spacing(3),
      overflowX: 'auto'
    },
    table: {
      minWidth: 650,
    },
  }));