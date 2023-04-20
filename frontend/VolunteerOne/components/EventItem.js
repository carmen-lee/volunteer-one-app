import React from "react";
import { withNavigation } from "@react-navigation/compat";
import PropTypes from "prop-types";
import {
  StyleSheet,
  Dimensions,
  Image,
  TouchableWithoutFeedback,
  Alert,
} from "react-native";
import { Block, Text, theme } from "galio-framework";

import { argonTheme } from "../constants";
import { Button } from "../components";
import { Avatar } from "@react-native-material/core";
import { useNavigation } from "@react-navigation/native";
const EventItem = (props) => {
  const navigation = useNavigation();
  console.log(props);
  const { full, style, imageStyle } = props;
  const handleRemoveItem = props.handleRemoveItem;
  const eventInfo = props.data;

  const imageStyles = [
    full ? styles.fullImage : styles.horizontalImage,
    imageStyle,
  ];
  const cardContainer = [styles.card, styles.shadow, style];
  const imgContainer = [
    styles.imageContainer,
    styles.horizontalStyles,

    styles.shadow,
  ];

  const confirmationAlert = (id) => {
    Alert.alert("Are you sure?", "", [
      {
        text: "Cancel",
        onPress: () => console.log("Cancel Pressed"),
        style: "cancel",
      },
      { text: "OK", onPress: () => handleRemoveItem(id) },
    ]);
  };
  return (
    <TouchableWithoutFeedback
      onPress={() =>
        navigation.navigate("ViewEvent", { eventID: eventInfo["id"] })
      }
    >
      <Block row={true} card style={cardContainer}>
        {/*================== profile image ==================*/}
        {/* TODO: link to profile  */}

        <Block>
          <Avatar style={imageStyles} />
        </Block>

        {/*================== username ==================*/}
        {/* TODO: link to profile */}
        {/* <TouchableWithoutFeedback onPress={() => navigation.navigate('Pro')}> */}
        {/* <TouchableWithoutFeedback onPress={() => navigation.push('Profile',  {theUser: item.username,})}>  */}
        <Block flex>
          <Text size={12} style={styles.cardTitle} bold>
            {eventInfo["organization"]}
          </Text>
          <Text size={12}> {eventInfo["name"]}</Text>
        </Block>
        {/* </TouchableWithoutFeedback> */}

        {/*================== buttons ==================*/}
        <Block row={true} style={styles.cardDescription}>
          {/*================== remove button ==================*/}
          <Button
            small
            style={{ backgroundColor: "grey" }}
            onPress={() => confirmationAlert(eventInfo["id"])}
          >
            Remove
          </Button>

          {/*================== options button (3 dots) ==================*/}
          {/* TODO: will need to change this touchable to a popup */}
          {/* <TouchableWithoutFeedback onPress={() => navigation.navigate('Pro')}>
              <Text size={12} style={styles.cardDescription} bold>•••</Text>
            </TouchableWithoutFeedback> */}
        </Block>
      </Block>
    </TouchableWithoutFeedback>
  );
};

EventItem.propTypes = {
  item: PropTypes.object,
  horizontal: PropTypes.bool,
  full: PropTypes.bool,
  ctaColor: PropTypes.string,
  imageStyle: PropTypes.any,
  following: PropTypes.bool,
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: theme.COLORS.WHITE,
    marginVertical: theme.SIZES.BASE / 2,
    minHeight: 96,
    justifyContent: "flex-start",
    alignItems: "center",
  },
  cardTitle: {},
  cardDescription: {
    // alignItems: 'center', // Centered vertically - 3 dots
    // flex:1,
    marginRight: 20,
    borderRadius: 3,
    elevation: 1,
    overflow: "hidden",
  },
  image: {},
  horizontalImage: {
    height: 50,
    width: 50,
    borderRadius: 62,
    margin: 20,
  },
  horizontalStyles: {
    borderTopRightRadius: 0,
    borderBottomRightRadius: 0,
  },
  verticalStyles: {
    borderBottomRightRadius: 0,
    borderBottomLeftRadius: 0,
  },
  fullImage: {
    height: 215,
  },
  shadow: {
    shadowColor: theme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 4,
    shadowOpacity: 0.1,
    elevation: 2,
  },
});

export default EventItem;
