package biz

import (
	"context"
	"crypto/md5"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/pkg/middleware/auth"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v5"
	"github.com/xuri/excelize/v2"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID                     int64
	Address                string
	Undo                   int64
	Password               string
	AddressTwo             string
	PrivateKey             string
	AddressThree           string
	WordThree              string
	PrivateKeyThree        string
	Last                   uint64
	LastBiw                uint64
	Amount                 uint64
	AmountBiw              uint64
	AmountUsdt             float64
	AmountUsdtOrigin       float64
	AmountUsdtGet          float64
	AmountRecommendUsdtGet float64
	MyTotalAmount          float64
	AmountFour             float64
	AmountFourGet          float64
	RecommendLevel         int64
	OutRate                int64
	Lock                   int64
	Vip                    int64
	VipAdmin               int64
	LockReward             int64
	CreatedAt              time.Time
	UpdatedAt              time.Time
	RecommendUserReward    int64
	RecommendUser          int64
	RecommendUserH         int64
	One                    string
	Two                    string
	Three                  string
	Four                   string
	Five                   string
	Six                    string
	Seven                  string
	AmountSelf             uint64
	IspayNew               float64
}

type Stake struct {
	ID        int64
	UserId    int64
	Status    int64
	Day       int64
	Amount    float64
	Reward    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Admin struct {
	ID       int64
	Password string
	Account  string
	Type     string
}

type AdminAuth struct {
	ID      int64
	AdminId int64
	AuthId  int64
}

type Auth struct {
	ID   int64
	Name string
	Path string
	Url  string
}

type UserInfo struct {
	ID               int64
	UserId           int64
	Vip              int64
	HistoryRecommend int64
	LockVip          int64
	UseVip           int64
	TeamCsdBalance   int64
}

type UserRecommendArea struct {
	ID            int64
	RecommendCode string
	Num           int64
}

type UserRecommend struct {
	ID            int64
	UserId        int64
	RecommendCode string
	Total         int64
	CreatedAt     time.Time
}

type UserBalanceRecord struct {
	ID           int64
	UserId       int64
	Balance      int64
	Amount       int64
	Type         string
	CoinType     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	BalanceNew   float64
	AmountNew    float64
	AmountNewTwo float64
}

type BalanceReward struct {
	ID             int64
	UserId         int64
	Status         int64
	Amount         int64
	SetDate        time.Time
	LastRewardDate time.Time
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type UserCurrentMonthRecommend struct {
	ID              int64
	UserId          int64
	RecommendUserId int64
	Date            time.Time
}

type Config struct {
	ID      int64
	KeyName string
	Name    string
	Value   string
}

type UserBalance struct {
	ID                     int64
	UserId                 int64
	BalanceUsdt            int64
	BalanceDhb             int64
	BalanceUsdtFloat       float64
	BalanceKsdtFloat       float64
	BalanceRawFloat        float64
	BalanceRawFloatNew     float64
	BalanceC               int64
	AreaTotalFloat         float64
	AreaTotalFloatTwo      float64
	AreaTotalFloatThree    float64
	RecommendTotalFloat    float64
	RecommendLevelFloat    float64
	RecommendTotalFloatTwo float64
	AllFloat               float64
	LocationTotalFloat     float64
}

type Withdraw struct {
	ID              int64
	UserId          int64
	Amount          int64
	RelAmount       int64
	AmountNew       float64
	RelAmountNew    float64
	BalanceRecordId int64
	Status          string
	Type            string
	Address         string
	CreatedAt       time.Time
}

type Trade struct {
	ID           int64
	UserId       int64
	AmountCsd    int64
	RelAmountCsd int64
	AmountHbs    int64
	CsdReward    int64
	RelAmountHbs int64
	Status       string
	CreatedAt    time.Time
}

type UserUseCase struct {
	repo                          UserRepo
	urRepo                        UserRecommendRepo
	configRepo                    ConfigRepo
	uiRepo                        UserInfoRepo
	ubRepo                        UserBalanceRepo
	locationRepo                  LocationRepo
	userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo
	tx                            Transaction
	log                           *log.Helper
}

type Reward struct {
	ID               int64
	UserId           int64
	Amount           int64
	AmountB          int64
	BalanceRecordId  int64
	AmountNew        float64
	AmountNewTwo     float64
	Type             string
	TypeRecordId     int64
	Status           int64
	Reason           string
	ReasonLocationId int64
	LocationType     string
	Address          string
	CreatedAt        time.Time
}

type Total struct {
	ID    int64
	One   float64
	Two   float64
	Three float64
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type UserArea struct {
	ID         int64
	UserId     int64
	Amount     int64
	SelfAmount int64
	Level      int64
}

type PriceChange struct {
	ID        int64
	Origin    int64
	Price     int64
	Status    int64
	CreatedAt time.Time
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

type BuyRecordFour struct {
	ID              int64
	UserId          int64
	Status          int64
	Amount          float64
	AmountGet       float64
	AmountGetPerDay float64
	LastUpdated     int64
	One             string
	Two             string
	Three           string
	Four            int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type BuyRecord struct {
	ID          int64
	UserId      int64
	Status      int64
	LastUpdated int64
	Amount      float64
	AmountGet   float64
	One         string
	Two         string
	Three       string
	Four        int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Good struct {
	ID     int64
	Amount uint64
	Name   string
	One    string
	Two    string
	Three  string
	Status uint64
}

type ConfigRepo interface {
	GetConfigByKeys(ctx context.Context, keys ...string) ([]*Config, error)
	GetConfigs(ctx context.Context) ([]*Config, error)
	UpdateConfig(ctx context.Context, id int64, value string) (bool, error)
	CreatePriceChangeConfig(ctx context.Context, origin int64, price int64) error
	UpdatePriceChangeStatus(ctx context.Context, id int64, status int64) (bool, error)
}

type UserBalanceRepo interface {
	GetPriceChangeConfig(ctx context.Context) (*PriceChange, error)
	GetStake(ctx context.Context) ([]*Stake, error)
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string, status string) (int64, error)
	WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string, status string) (int64, error)
	RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error)
	RecommendTeamReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, locationId int64, recommendNum int64, status string) (int64, error)
	RecommendRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64, userIdTwo int64) (int64, error)
	LocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, stop string, currentMaxNew int64, feeRate int64) (int64, error)
	RecommendLocationRewardBiw(ctx context.Context, userId int64, rewardAmount int64, recommendNum int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	PriceChange(ctx context.Context, userId int64, rewardAmount int64, up string) error
	AreaRewardBiw(ctx context.Context, userId int64, rewardAmount int64, tmpCurrentReward int64, areaType int64, stop string, tmpMaxNew int64, feeRate int64) (int64, error)
	FourRewardBiw(ctx context.Context, userId int64, rewardAmount int64, num int64) (int64, error)
	FourRewardYes(ctx context.Context, rewardAmount int64) error
	ExchangeBiw(ctx context.Context, userId int64, currentMaxNew int64, feeRate int64) (int64, error)
	GetRewardFourYes(ctx context.Context) (*Reward, error)
	SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error
	SystemReward(ctx context.Context, amount int64, locationId int64) error
	SystemDailyReward(ctx context.Context, amount int64, locationId int64) error
	GetSystemYesterdayDailyReward(ctx context.Context, day int) (*Reward, error)
	GetTotal(ctx context.Context) (*Total, error)
	GetSystemYesterdayLocationReward(ctx context.Context, day int) ([]*UserBalanceRecord, error)
	GetRewardYes(ctx context.Context) ([]*Reward, error)
	SystemFee(ctx context.Context, amount int64, locationId int64) error
	UserFee(ctx context.Context, userId int64, amount int64) (int64, error)
	UserDailyFee(ctx context.Context, userId int64, amount int64, status string) (int64, error)
	UserDailyRecommendArea(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, status string) (int64, error)
	RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error)
	RecommendWithdrawTopReward(ctx context.Context, userId int64, amount int64, locationId int64, vip int64, status string) (int64, error)
	NormalRecommendReward(ctx context.Context, userId int64, rewardAmount int64, rewardAmount2 int64, locationId int64, status string, status2 string, type1 string, reason string) (int64, error)
	NewNormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	NormalRecommendReward2(ctx context.Context, userId int64, rewardAmount int64, locationId int64, type1 string, reason string) (int64, error)
	NormalReward3(ctx context.Context, userId int64, rewardAmount int64, rewardAmount2 int64, locationId int64, status string, status2 string) (int64, error)
	NormalReward4(ctx context.Context, userId int64, rewardAmount int64, locationId int64) (int64, error)
	NormalRecommendTopReward(ctx context.Context, userId int64, amount int64, locationId int64, reasonId int64, status string) (int64, error)
	NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64, status string) (int64, error)
	NormalWithdrawRecommendTopReward(ctx context.Context, userId int64, amount int64, locationId int64, reasonId int64, status string) (int64, error)
	Deposit(ctx context.Context, userId int64, amount int64, dhbAmount int64) (int64, error)
	DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error)
	DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error)
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserRewardByUserId(ctx context.Context, userId int64) ([]*Reward, error)
	GetUserRewards(ctx context.Context, b *Pagination, userId int64, reason string) ([]*Reward, error, int64)
	GetUserBuy(ctx context.Context, b *Pagination, userId int64) ([]*BuyRecord, error, int64)
	GetUserBuyTwo(ctx context.Context, b *Pagination, userId int64) ([]*BuyRecord, error, int64)
	GetUserBuyThree(ctx context.Context, b *Pagination, userId int64) ([]*BuyRecord, error, int64)
	GetGoods(ctx context.Context) ([]*Good, error)
	GetGoodsTwo(ctx context.Context) ([]*Good, error)
	GetGoodsThree(ctx context.Context) ([]*Good, error)
	GetGoodsOnline(ctx context.Context) ([]*Good, error)
	GetGoodsOnlineTwo(ctx context.Context) ([]*Good, error)
	GetGoodsOnlineThree(ctx context.Context) ([]*Good, error)
	GetGoodsPage(ctx context.Context, b *Pagination) ([]*Good, error, int64)
	GetGoodsPageTwo(ctx context.Context, b *Pagination) ([]*Good, error, int64)
	GetGoodsPageThree(ctx context.Context, b *Pagination) ([]*Good, error, int64)
	GetUserBuyByUserId(ctx context.Context, userId int64) ([]*BuyRecord, error)
	GetUserBuyById(id int64) (*BuyRecord, error)
	GetUserBuyFour(ctx context.Context, b *Pagination, userId int64) ([]*BuyRecordFour, error, int64)
	GetUserRewardsLastMonthFee(ctx context.Context) ([]*Reward, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceLockByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserBalance, error)
	GetUserBalanceUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceLockUsdtTotal(ctx context.Context) (int64, error)
	GetUserLocationNewCurrentMaxNew(ctx context.Context) (int64, error)
	GetUserLocationNewCurrentMax(ctx context.Context) (int64, error)
	GetUserLocationNewCurrent(ctx context.Context) (int64, error)
	GetUserBalanceDHBTotal(ctx context.Context) (int64, error)
	GreateWithdraw(ctx context.Context, userId int64, amount int64, coinType string) (*Withdraw, error)
	WithdrawUsdt(ctx context.Context, userId int64, amount int64) error
	WithdrawDhb(ctx context.Context, userId int64, amount int64) error
	GetWithdrawByUserId(ctx context.Context, userId int64) ([]*Withdraw, error)
	GetWithdraws(ctx context.Context, b *Pagination, userId int64, withdrawType string) ([]*Withdraw, error, int64)
	GetWithdrawPassOrRewarded(ctx context.Context) ([]*Withdraw, error)
	GetWithdrawByUserIdsMap(ctx context.Context, userIds []int64) (map[int64][]*Withdraw, error)
	GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error)
	GetTradeOk(ctx context.Context) (*Trade, error)
	GetTradeOkkCsd(ctx context.Context) (int64, error)
	GetTradeOkkHbs(ctx context.Context) (int64, error)
	UpdateWithdraw(ctx context.Context, id int64, status string) (*Withdraw, error)
	GetWithdrawById(ctx context.Context, id int64) (*Withdraw, error)
	GetWithdrawNotDeal(ctx context.Context) ([]*Withdraw, error)
	GetWithdrawByUserIds(ctx context.Context, userIds []int64) ([]*Withdraw, error)
	GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalTwo(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalThree(ctx context.Context) (int64, error)
	GetUserBalanceRecordCsdTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordHbsTotal(ctx context.Context) (int64, error)
	GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserRewardLocationTotalToday(ctx context.Context, reason string) (int64, error)
	GetSystemWithdrawUsdtFeeTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawDhbTotalToday(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error)
	GetUserWithdrawDhbTotal(ctx context.Context) (int64, error)
	GetUserWithdrawUsdtTotalByUserIds(ctx context.Context, userIds []int64) (int64, error)
	GetUserRewardUsdtTotal(ctx context.Context) (int64, error)
	GetUserRewardBalanceRewardTotal(ctx context.Context) (int64, error)
	GetBalanceRewardTotal(ctx context.Context) (int64, error)
	GetSystemRewardUsdtTotal(ctx context.Context) (int64, error)
	UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*Withdraw, error)
	GetUserRewardRecommendSort(ctx context.Context) ([]*UserSortRecommendReward, error)
	UpdateBalance(ctx context.Context, userId int64, amount int64) (bool, error)
	UpdateTrade(ctx context.Context, id int64, status string) (*Trade, error)
	GetTradeNotDeal(ctx context.Context) ([]*Trade, error)

	UpdateWithdrawPass(ctx context.Context, id int64) (*Withdraw, error)
	UserDailyBalanceReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, amountDhb int64, status string) (int64, error)
	GetBalanceRewardCurrent(ctx context.Context, now time.Time) ([]*BalanceReward, error)
	GetUserTrades(ctx context.Context, b *Pagination, userId int64) ([]*Trade, error, int64)
	UserDailyLocationReward(ctx context.Context, userId int64, rewardAmount int64, amount int64, coinAmount int64, status string, locationId int64) (int64, error)
	DepositLastNew(ctx context.Context, userId int64, lastAmount int64) (int64, error)
	DepositLastNew2(ctx context.Context, userId int64, lastAmount int64) (int64, error)
	DepositLastNewDhb(ctx context.Context, userId int64, lastCoinAmount int64) error
	DepositLastNewCsd(ctx context.Context, userId int64, lastCoinAmount int64, tmpRecommendUserIdsInt []int64) error
	UpdateBalanceRewardLastRewardDate(ctx context.Context, id int64) error
	UpdateLocationAgain(ctx context.Context, locations []*LocationNew) error
	UpdateLocationAgain2(ctx context.Context, locations []*LocationNew) error
	LocationNewDailyReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error)
	WithdrawNewRewardRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	WithdrawNewRewardTeamRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	WithdrawNewRewardSecondRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	WithdrawNewRewardLevelRecommend(ctx context.Context, userId int64, amount int64, amountB int64, locationId int64, tmpRecommendUserIdsInt []int64) (int64, error)
	UpdateLocationNewMax(ctx context.Context, userId int64, amount int64) (int64, error)
	GetAllUsersB(ctx context.Context) ([]*User, error)
}

type UserRecommendRepo interface {
	GetUserRecommendByUserId(ctx context.Context, userId int64) (*UserRecommend, error)
	CreateUserRecommend(ctx context.Context, u *User, recommendUser *UserRecommend) (*UserRecommend, error)
	GetUserRecommendByCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommendLikeCode(ctx context.Context, code string) ([]*UserRecommend, error)
	GetUserRecommends(ctx context.Context) ([]*UserRecommend, error)
	GetUserRecommendsFour(ctx context.Context) ([]*UserRecommend, error)
	CreateUserRecommendArea(ctx context.Context, recommendAreas []*UserRecommendArea) (bool, error)
	GetUserRecommendLowAreas(ctx context.Context) ([]*UserRecommendArea, error)
	UpdateUserAreaAmount(ctx context.Context, userId int64, amount int64) (bool, error)
	UpdateUserAreaSelfAmount(ctx context.Context, userId int64, amount int64) (bool, error)
	UpdateUserAreaLevel(ctx context.Context, userId int64, level int64) (bool, error)
	UpdateUserRecommendTotal(ctx context.Context, userId int64, total int64) error
	UpdateUserAreaLevelUp(ctx context.Context, userId int64, level int64) (bool, error)
	GetUserAreas(ctx context.Context, userIds []int64) ([]*UserArea, error)
	GetUserArea(ctx context.Context, userId int64) (*UserArea, error)
	CreateUserArea(ctx context.Context, u *User) (bool, error)
}

type UserCurrentMonthRecommendRepo interface {
	GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *Pagination, userId int64) ([]*UserCurrentMonthRecommend, error, int64)
	CreateUserCurrentMonthRecommend(ctx context.Context, u *UserCurrentMonthRecommend) (*UserCurrentMonthRecommend, error)
	GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error)
	GetUserLastMonthRecommend(ctx context.Context) ([]int64, error)
}

type UserInfoRepo interface {
	UpdateUserNewTwoNewTwo(ctx context.Context, userId int64, amount uint64, amountIspay float64, one, two, three string, four int64) error
	UpdateUserNewTwoNewTwoTwo(ctx context.Context, userId int64, amount uint64) error
	UpdateUserNewTwoNewTwoNew(ctx context.Context, userId int64, amount uint64, one, two, three string, four int64) error
	GetAllBuyRecord(ctx context.Context) ([]*BuyRecord, error)
	GetBuyRecordMap(ctx context.Context, userIds []int64) (map[int64][]*BuyRecord, error)
	GetBuyRecordingMap(ctx context.Context, userIds []int64) (map[int64][]*BuyRecord, error)
	UpdateUserRewardStakeReomve(ctx context.Context, userId int64, amountUsdt float64, stakeId int64) (int64, error)
	UpdateUserRewardStake(ctx context.Context, userId int64, amountUsdt float64, stakeId int64) (int64, error)
	UpdateUserRewardNew(ctx context.Context, id, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool) (int64, error)
	UpdateUserRewardNewFour(ctx context.Context, userId int64, amountUsdt float64) (int64, error)
	UpdateUserRewardRecommendNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64, address string) (int64, error)
	UpdateUserReward(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool) (int64, error)
	UpdateUserRewardRecommend(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, address string) (int64, error)
	UpdateUserRewardArea(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, tmpLevel, stop bool, level, i int64, address string) (int64, error)
	UpdateUserRewardAreaTwo(ctx context.Context, userId int64, amountUsdt float64, stop bool) (int64, error)
	UpdateUserRewardRecommendUserGet(ctx context.Context, userId int64, amountUsdt float64, enough bool, amount float64) error
	UpdateUserMyTotalAmount(ctx context.Context, userId int64, amountUsdt float64) error
	UpdateUserAmountSelf(ctx context.Context, userId int64, amountUsdt uint64) error
	UpdateUserMyTotalAmountAdd(ctx context.Context, userId int64, amountUsdt, myTotal float64) error
	UpdateUserMyTotalAmountSub(ctx context.Context, userId int64, amountUsdt float64) error
	GetBuyRecord(ctx context.Context, day int) ([]*BuyRecord, error)
	GetBuyRecordFour(ctx context.Context) ([]*BuyRecordFour, error)
	UpdateTotalOne(ctx context.Context, amountUsdt float64) error
	UpdateUserNewTwoNewThree(ctx context.Context, userId int64, amount uint64, last uint64, coinType string) error
	UpdateUserIspay(ctx context.Context, userId int64, amount uint64) error
	UpdateUserUsdtFloat(ctx context.Context, userId int64, amount float64, last float64, coinType string) error
	UpdateUserRecommendLevel(ctx context.Context, userId int64, level uint64) error
	UpdateUserRecommendLevel2(ctx context.Context, userId int64, level uint64) error
	UpdateUserPass(ctx context.Context, userId int64, pass string) error
	UpdateUserAmountFour(ctx context.Context, userId int64, amountFour float64, do bool) error
	UpdateUserLast(ctx context.Context, userId int64, coinType string) error
	CreateUserInfo(ctx context.Context, u *User) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	UpdateUserPassword(ctx context.Context, userId int64, password string) (*User, error)
	UpdateUserAddress(ctx context.Context, address string, addressTwo string) error
	UpdateUserInfo(ctx context.Context, u *UserInfo) (*UserInfo, error)
	UpdateUserInfo2(ctx context.Context, u *UserInfo) (*UserInfo, error)
	UpdateUserInfoVip(ctx context.Context, userId, vip int64) (*UserInfo, error)
	GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*UserInfo, error)
	GetUserInfosByVipAndLockVip(ctx context.Context) ([]*UserInfo, error)
	UpdateUserRewardAreaNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, level, i int64, address string) (int64, error)
	UpdateUserRewardTotalOneNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64) (int64, error)
	UpdateUserRewardTotalTwoNew(ctx context.Context, userId int64, amountUsdt float64, amountUsdtTotal float64, stop bool, i int64) (int64, error)
	UpdateUserRewardTotalOver(ctx context.Context) error
	UpdateUserRewardRecommend2(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string) error
	UpdateUserRewardDailyLocation(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool) error
	UpdateUserRewardDailyIspay(ctx context.Context, id, userId int64, raw, rawOrigin, usdtOrigin float64, stop bool) error
	UpdateUserSubBuyRecord(ctx context.Context, id, userId int64, amountOrigin float64) error
	UpdateUserRewardAreaOne(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string, i, cl int64, two bool) error
	UpdateUserRewardRecommendNewTwo(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool, address string, i int64) error
	UpdateUserRewardAllNew(ctx context.Context, id, userId int64, usdt, raw, usdtOrigin float64, amountOrigin float64, stop bool) error
	UpdateUserNewNewNewFour(ctx context.Context, userId int64, amount uint64, amountIspay, amountIspayPerDay float64, one, two, three string) error
}

type UserRepo interface {
	UpdateGoods(ctx context.Context, id, status uint64) error
	UpdateGoodsTwo(ctx context.Context, id, status uint64) error
	UpdateGoodsThree(ctx context.Context, id, status uint64) error
	CreateGoods(ctx context.Context, one, name, picName, three string, amount uint64) error
	CreateGoodsTwo(ctx context.Context, one, name, picName, three string, amount uint64) error
	CreateGoodsThree(ctx context.Context, one, name, picName, three string, amount uint64) error
	GetRewardYes(ctx context.Context) ([]*Reward, error)
	GetUsersNewTwo(ctx context.Context) ([]*User, error)
	GetUserById(ctx context.Context, Id int64) (*User, error)
	UndoUser(ctx context.Context, userId int64, undo int64) (bool, error)
	LockUser(ctx context.Context, userId int64, lock int64) (bool, error)
	LockUserReward(ctx context.Context, userId int64, lock int64) (bool, error)
	GetAdminByAccount(ctx context.Context, account string, password string) (*Admin, error)
	GetAdminById(ctx context.Context, id int64) (*Admin, error)
	GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserByAddressTwo(ctx context.Context, address string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	CreateAdmin(ctx context.Context, a *Admin) (*Admin, error)
	GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetAdmins(ctx context.Context) ([]*Admin, error)
	GetUsers(ctx context.Context, b *Pagination, address string, isLocation bool, vip int64) ([]*User, error, int64)
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetAllUsersOrderAmountBiw(ctx context.Context) ([]*User, error)
	GetAllUsersRecommendOrder(ctx context.Context) ([]*User, error)
	GetAllUserInfos(ctx context.Context) ([]*UserInfo, error)
	GetAllUserBalance(ctx context.Context) ([]*UserBalance, error)
	GetBuyRecord(ctx context.Context, day int) ([]*BuyRecord, error)
	GetUserCount(ctx context.Context) (int64, error)
	GetUserByUserIdsTwo(ctx context.Context, userIds []int64) (map[int64]*User, error)
	UpdateUserVip(ctx context.Context, userId int64, vip int64) error
	GetUserCountToday(ctx context.Context) (int64, error)
	CreateAdminAuth(ctx context.Context, adminId int64, authId int64) (bool, error)
	DeleteAdminAuth(ctx context.Context, adminId int64, authId int64) (bool, error)
	GetAuths(ctx context.Context) ([]*Auth, error)
	GetAuthByIds(ctx context.Context, ids ...int64) (map[int64]*Auth, error)
	GetAdminAuth(ctx context.Context, adminId int64) ([]*AdminAuth, error)
	UpdateAdminPassword(ctx context.Context, account string, password string) (*Admin, error)
}

func NewUserUseCase(repo UserRepo, tx Transaction, configRepo ConfigRepo, uiRepo UserInfoRepo, urRepo UserRecommendRepo, locationRepo LocationRepo, userCurrentMonthRecommendRepo UserCurrentMonthRecommendRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:                          repo,
		tx:                            tx,
		configRepo:                    configRepo,
		locationRepo:                  locationRepo,
		userCurrentMonthRecommendRepo: userCurrentMonthRecommendRepo,
		uiRepo:                        uiRepo,
		urRepo:                        urRepo,
		ubRepo:                        ubRepo,
		log:                           log.NewHelper(logger),
	}
}

func (uuc *UserUseCase) GetUsersNewTwo(ctx context.Context) ([]*User, error) {
	return uuc.repo.GetUsersNewTwo(ctx)
}

func (uuc *UserUseCase) GetUserByAddress(ctx context.Context, Addresses ...string) (map[string]*User, error) {
	return uuc.repo.GetUserByAddresses(ctx, Addresses...)
}

func (uuc *UserUseCase) GetUserByAddressTwo(ctx context.Context, Address string) (*User, error) {
	return uuc.repo.GetUserByAddressTwo(ctx, Address)
}

func (uuc *UserUseCase) GetbPriceConfig(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "b_price")
}

func (uuc *UserUseCase) GetExistUserByAddressOrCreate(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user *User
	)
	return user, nil
}

func (uuc *UserUseCase) UserInfo(ctx context.Context, user *User) (*v1.UserInfoReply, error) {
	return &v1.UserInfoReply{}, nil
}

func (uuc *UserUseCase) RewardList(ctx context.Context, req *v1.RewardListRequest, user *User) (*v1.RewardListReply, error) {
	res := &v1.RewardListReply{
		Rewards: make([]*v1.RewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) RecommendRewardList(ctx context.Context, user *User) (*v1.RecommendRewardListReply, error) {
	res := &v1.RecommendRewardListReply{
		Rewards: make([]*v1.RecommendRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) FeeRewardList(ctx context.Context, user *User) (*v1.FeeRewardListReply, error) {
	res := &v1.FeeRewardListReply{
		Rewards: make([]*v1.FeeRewardListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) WithdrawList(ctx context.Context, user *User) (*v1.WithdrawListReply, error) {
	res := &v1.WithdrawListReply{
		Withdraw: make([]*v1.WithdrawListReply_List, 0),
	}

	return res, nil
}

func (uuc *UserUseCase) Withdraw(ctx context.Context, req *v1.WithdrawRequest, user *User) (*v1.WithdrawReply, error) {
	return &v1.WithdrawReply{
		Status: "ok",
	}, nil
}

func (uuc *UserUseCase) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	var (
		userSearch  *User
		userId      int64 = 0
		userRewards []*Reward
		users       map[int64]*User
		userIdsMap  map[int64]int64
		userIds     []int64
		err         error
		count       int64
	)
	res := &v1.AdminRewardListReply{
		Rewards: make([]*v1.AdminRewardListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userRewards, err, count = uuc.ubRepo.GetUserRewards(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, req.Reason)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		tmpReason := vUserReward.Reason

		tmpLevel := vUserReward.BalanceRecordId
		tmpNum := vUserReward.TypeRecordId
		amountNew := fmt.Sprintf("%.4f", vUserReward.AmountNew)

		res.Rewards = append(res.Rewards, &v1.AdminRewardListReply_List{
			CreatedAt:  vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:     amountNew,
			AmountNana: fmt.Sprintf("%.4f", vUserReward.AmountNewTwo),
			Address:    tmpUser,
			Reason:     tmpReason,
			Num:        tmpNum,   // 代数
			Level:      tmpLevel, // 级别
			AddressTwo: vUserReward.Address,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminTradeList(ctx context.Context, req *v1.AdminTradeListRequest) (*v1.AdminTradeListReply, error) {
	var (
		userSearch *User
		userId     int64 = 0
		userTrades []*Trade
		users      map[int64]*User
		userIdsMap map[int64]int64
		userIds    []int64
		err        error
		count      int64
	)
	res := &v1.AdminTradeListReply{
		Trades: make([]*v1.AdminTradeListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userTrades, err, count = uuc.ubRepo.GetUserTrades(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userTrades {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userTrades {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		res.Trades = append(res.Trades, &v1.AdminTradeListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			AmountCsd: fmt.Sprintf("%.2f", float64(vUserReward.AmountCsd)/float64(100000)),
			AmountHbs: fmt.Sprintf("%.2f", float64(vUserReward.AmountHbs)/float64(100000)),
			Address:   tmpUser,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {
	var (
		users        []*User
		userIds      []int64
		userBalances map[int64]*UserBalance
		count        int64
		err          error
	)

	res := &v1.AdminUserListReply{
		Users: make([]*v1.AdminUserListReply_UserList, 0),
	}

	users, err, count = uuc.repo.GetUsers(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, req.Address, false, 0)
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}

	userBalances, err = uuc.ubRepo.GetUserBalanceByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		myLowUser         map[int64][]*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
	)

	myLowUser = make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)

	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		usersAll []*User
		usersMap map[int64]*User
	)
	usersAll, err = uuc.repo.GetAllUsers(ctx)
	if nil == usersAll {
		return nil, nil
	}
	usersMap = make(map[int64]*User, 0)

	for _, vUsers := range usersAll {
		usersMap[vUsers.ID] = vUsers
	}

	t := time.Date(2026, 2, 18, 14, 0, 0, 0, time.UTC)
	for _, vUsers := range users {
		if _, ok := userBalances[vUsers.ID]; !ok {
			continue
		}

		var (
			userBuysFour []*BuyRecordFour
		)
		userBuysFour, err, _ = uuc.ubRepo.GetUserBuyFour(ctx, nil, vUsers.ID)
		if nil != err {
			continue
		}

		var tmpFourAmountPerDay float64
		for _, vBuyFour := range userBuysFour {
			tmpFourAmountPerDay += vBuyFour.AmountGetPerDay
		}

		var (
			userBuys []*BuyRecord
		)
		userBuys, err = uuc.ubRepo.GetUserBuyByUserId(ctx, vUsers.ID)
		if nil != err {
			continue
		}

		tmpAll := float64(0)
		tmpGet := float64(0)
		tmpGetSub := float64(0)
		tmpCurrentUsdtAmount := uint64(0)
		if 0 < vUsers.Amount {
			tmpCurrentUsdtAmount = vUsers.Amount
		}
		for _, vBuy := range userBuys {
			num := 2.5
			if vBuy.CreatedAt.After(t) {
				amountB := uint64(vBuy.Amount)
				if 4999 <= amountB && 15001 > amountB {
					num = 3
				} else if 29999 <= amountB && 50001 > amountB {
					num = 3.5
				} else if 99999 <= amountB && 150001 > amountB {
					num = 4
				}
			}
			tmpAll += vBuy.Amount * num
			tmpGet += vBuy.AmountGet
		}

		if tmpAll > tmpGet {
			tmpGetSub = tmpAll - tmpGet
		}

		tmpMyRecommendUserIdsLen := int64(0)
		tmpMax := uint64(0)
		tmpAreaMin := uint64(0)
		tmpMaxId := int64(0)
		if _, ok := myLowUser[vUsers.ID]; ok {
			tmpMyRecommendUserIdsLen = int64(len(myLowUser[vUsers.ID]))

			for _, vV := range myLowUser[vUsers.ID] {
				if _, ok2 := usersMap[vV.UserId]; ok2 {
					if tmpMax < uint64(usersMap[vV.UserId].MyTotalAmount)+usersMap[vV.UserId].AmountSelf {
						tmpMax = uint64(usersMap[vV.UserId].MyTotalAmount) + usersMap[vV.UserId].AmountSelf
						tmpMaxId = vV.ID
					}
				}
			}

			if 0 < tmpMaxId {
				for _, vMyLowUser := range myLowUser[vUsers.ID] {
					if _, ok2 := usersMap[vMyLowUser.UserId]; !ok2 {
						continue
					}

					if tmpMaxId != vMyLowUser.ID {
						tmpAreaMin += uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].AmountSelf
					}
				}
			}
		}

		currentLevel := int64(0)
		if 1500000 <= tmpAreaMin {
			currentLevel = 5
		} else if 500000 <= tmpAreaMin {
			currentLevel = 4
		} else if 150000 <= tmpAreaMin {
			currentLevel = 3
		} else if 50000 <= tmpAreaMin {
			currentLevel = 2
		} else if 10000 <= tmpAreaMin {
			currentLevel = 1
		}

		if 0 < vUsers.VipAdmin {
			currentLevel = vUsers.VipAdmin
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)

		addressMyRecommend := ""
		if _, ok := userRecommendsMap[vUsers.ID]; ok {
			userRecommend = userRecommendsMap[vUsers.ID]

			if nil != userRecommend && "" != userRecommend.RecommendCode {
				var (
					tmpRecommendUserIds   []string
					myUserRecommendUserId int64
				)
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
				if 2 <= len(tmpRecommendUserIds) {
					myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
				}

				if 0 < myUserRecommendUserId {
					if _, ok2 := usersMap[myUserRecommendUserId]; ok2 {
						addressMyRecommend = usersMap[myUserRecommendUserId].Address
					}
				}
			}
		}

		res.Users = append(res.Users, &v1.AdminUserListReply_UserList{
			UserId:             vUsers.ID,
			CreatedAt:          vUsers.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:            vUsers.Address,
			BalanceUsdt:        fmt.Sprintf("%.2f", userBalances[vUsers.ID].BalanceUsdtFloat),
			BalanceDhb:         fmt.Sprintf("%.2f", userBalances[vUsers.ID].BalanceRawFloat),
			BAmount:            fmt.Sprintf("%.2f", userBalances[vUsers.ID].BalanceRawFloatNew),
			BAmountTwo:         fmt.Sprintf("%.4f", vUsers.IspayNew),
			PerDayAmount:       fmt.Sprintf("%.4f", tmpFourAmountPerDay),
			Vip:                currentLevel,
			Out:                vUsers.OutRate,
			HistoryRecommend:   tmpMyRecommendUserIdsLen,
			AreaTotal:          vUsers.MyTotalAmount,
			AreaMax:            float64(tmpMax),
			AreaMin:            float64(tmpAreaMin),
			AmountUsdtGet:      fmt.Sprintf("%.2f", tmpGet),
			AmountUsdtCurrent:  fmt.Sprintf("%.d", tmpCurrentUsdtAmount),
			AmountUsdtTwo:      fmt.Sprintf("%.2f", tmpGetSub),
			Lock:               vUsers.Lock,
			LockReward:         vUsers.LockReward,
			MyRecommendAddress: addressMyRecommend,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminBuyFourList(ctx context.Context, req *v1.AdminBuyListRequest) (*v1.AdminBuyListReply, error) {

	var (
		userSearch  *User
		userId      int64 = 0
		userRewards []*BuyRecordFour
		users       map[int64]*User
		userIdsMap  map[int64]int64
		userIds     []int64
		err         error
		count       int64
		goods       []*Good
		goodsMap    map[int64]*Good
	)
	res := &v1.AdminBuyListReply{
		Rewards: make([]*v1.AdminBuyListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userRewards, err, count = uuc.ubRepo.GetUserBuyFour(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	goods, err = uuc.ubRepo.GetGoods(ctx)
	if nil != err {
		return nil, err
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		oneTmp := ""
		if "1" != vUserReward.One {
			oneTmp = vUserReward.One
		}
		twoTmp := ""
		if "1" != vUserReward.One {
			twoTmp = vUserReward.Two
		}
		threeTmp := ""
		if "1" != vUserReward.One {
			threeTmp = vUserReward.Three
		}
		fourTmp := ""
		if 0 != vUserReward.Four {
			if _, ok := goodsMap[vUserReward.Four]; ok {
				fourTmp = goodsMap[vUserReward.Four].Name
			}
		}

		res.Rewards = append(res.Rewards, &v1.AdminBuyListReply_List{
			CreatedAt:    vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:       strconv.FormatInt(vUserReward.Four, 10),
			Address:      tmpUser,
			Id:           vUserReward.ID,
			One:          fourTmp,
			Two:          oneTmp,
			Three:        twoTmp,
			Four:         threeTmp,
			AmountTotal:  fmt.Sprintf("%.4f", vUserReward.Amount),
			AmountGet:    fmt.Sprintf("%.4f", vUserReward.AmountGet),
			PerDayAmount: fmt.Sprintf("%.4f", vUserReward.AmountGetPerDay),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminBuyList(ctx context.Context, req *v1.AdminBuyListRequest) (*v1.AdminBuyListReply, error) {

	var (
		userSearch  *User
		userId      int64 = 0
		userRewards []*BuyRecord
		users       map[int64]*User
		userIdsMap  map[int64]int64
		userIds     []int64
		err         error
		count       int64
		goods       []*Good
		goodsMap    map[int64]*Good
	)
	res := &v1.AdminBuyListReply{
		Rewards: make([]*v1.AdminBuyListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userRewards, err, count = uuc.ubRepo.GetUserBuy(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	goods, err = uuc.ubRepo.GetGoods(ctx)
	if nil != err {
		return nil, err
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		oneTmp := ""
		if "1" != vUserReward.One {
			oneTmp = vUserReward.One
		}
		twoTmp := ""
		if "1" != vUserReward.One {
			twoTmp = vUserReward.Two
		}
		threeTmp := ""
		if "1" != vUserReward.One {
			threeTmp = vUserReward.Three
		}
		fourTmp := ""
		if 0 != vUserReward.Four {
			if _, ok := goodsMap[vUserReward.Four]; ok {
				fourTmp = goodsMap[vUserReward.Four].Name
			}
		}

		res.Rewards = append(res.Rewards, &v1.AdminBuyListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", vUserReward.Amount),
			Address:   tmpUser,
			Id:        vUserReward.ID,
			One:       fourTmp,
			Two:       oneTmp,
			Three:     twoTmp,
			Four:      threeTmp,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminBuyListTwo(ctx context.Context, req *v1.AdminBuyListRequest) (*v1.AdminBuyListReply, error) {

	var (
		userSearch  *User
		userId      int64 = 0
		userRewards []*BuyRecord
		users       map[int64]*User
		userIdsMap  map[int64]int64
		userIds     []int64
		err         error
		count       int64
		goods       []*Good
		goodsMap    map[int64]*Good
	)
	res := &v1.AdminBuyListReply{
		Rewards: make([]*v1.AdminBuyListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userRewards, err, count = uuc.ubRepo.GetUserBuyTwo(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	goods, err = uuc.ubRepo.GetGoodsTwo(ctx)
	if nil != err {
		return nil, err
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		oneTmp := ""
		if "1" != vUserReward.One {
			oneTmp = vUserReward.One
		}
		twoTmp := ""
		if "1" != vUserReward.One {
			twoTmp = vUserReward.Two
		}
		threeTmp := ""
		if "1" != vUserReward.One {
			threeTmp = vUserReward.Three
		}
		fourTmp := ""
		if 0 != vUserReward.Four {
			if _, ok := goodsMap[vUserReward.Four]; ok {
				fourTmp = goodsMap[vUserReward.Four].Name
			}
		}

		res.Rewards = append(res.Rewards, &v1.AdminBuyListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", vUserReward.Amount),
			Address:   tmpUser,
			Id:        vUserReward.ID,
			One:       fourTmp,
			Two:       oneTmp,
			Three:     twoTmp,
			Four:      threeTmp,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminBuyListThree(ctx context.Context, req *v1.AdminBuyListRequest) (*v1.AdminBuyListReply, error) {

	var (
		userSearch  *User
		userId      int64 = 0
		userRewards []*BuyRecord
		users       map[int64]*User
		userIdsMap  map[int64]int64
		userIds     []int64
		err         error
		count       int64
		goods       []*Good
		goodsMap    map[int64]*Good
	)
	res := &v1.AdminBuyListReply{
		Rewards: make([]*v1.AdminBuyListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	userRewards, err, count = uuc.ubRepo.GetUserBuyThree(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vUserReward := range userRewards {
		userIdsMap[vUserReward.UserId] = vUserReward.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	goods, err = uuc.ubRepo.GetGoodsThree(ctx)
	if nil != err {
		return nil, err
	}
	goodsMap = make(map[int64]*Good, 0)
	for _, v := range goods {
		goodsMap[v.ID] = v
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	for _, vUserReward := range userRewards {
		tmpUser := ""
		if nil != users {
			if _, ok := users[vUserReward.UserId]; ok {
				tmpUser = users[vUserReward.UserId].Address
			}
		}

		oneTmp := ""
		if "1" != vUserReward.One {
			oneTmp = vUserReward.One
		}
		twoTmp := ""
		if "1" != vUserReward.One {
			twoTmp = vUserReward.Two
		}
		threeTmp := ""
		if "1" != vUserReward.One {
			threeTmp = vUserReward.Three
		}
		fourTmp := ""
		if 0 != vUserReward.Four {
			if _, ok := goodsMap[vUserReward.Four]; ok {
				fourTmp = goodsMap[vUserReward.Four].Name
			}
		}

		res.Rewards = append(res.Rewards, &v1.AdminBuyListReply_List{
			CreatedAt: vUserReward.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", vUserReward.Amount),
			Address:   tmpUser,
			Id:        vUserReward.ID,
			One:       fourTmp,
			Two:       oneTmp,
			Three:     twoTmp,
			Four:      threeTmp,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminGoodList(ctx context.Context, req *v1.AdminGoodListRequest) (*v1.AdminGoodListReply, error) {

	var (
		goods []*Good
		count int64
		err   error
	)
	res := &v1.AdminGoodListReply{
		Goods: make([]*v1.AdminGoodListReply_List, 0),
	}

	goods, err, count = uuc.ubRepo.GetGoodsPage(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, v := range goods {
		res.Goods = append(res.Goods, &v1.AdminGoodListReply_List{
			Id:     v.ID,
			Name:   v.Name,
			One:    v.One,
			Two:    v.Two,
			Amount: v.Amount,
			Three:  v.Three,
			Status: v.Status,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminGoodListTwo(ctx context.Context, req *v1.AdminGoodListRequest) (*v1.AdminGoodListReply, error) {

	var (
		goods []*Good
		count int64
		err   error
	)
	res := &v1.AdminGoodListReply{
		Goods: make([]*v1.AdminGoodListReply_List, 0),
	}

	goods, err, count = uuc.ubRepo.GetGoodsPageTwo(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, v := range goods {
		res.Goods = append(res.Goods, &v1.AdminGoodListReply_List{
			Id:     v.ID,
			Name:   v.Name,
			One:    v.One,
			Two:    v.Two,
			Amount: v.Amount,
			Three:  v.Three,
			Status: v.Status,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminGoodListThree(ctx context.Context, req *v1.AdminGoodListRequest) (*v1.AdminGoodListReply, error) {

	var (
		goods []*Good
		count int64
		err   error
	)
	res := &v1.AdminGoodListReply{
		Goods: make([]*v1.AdminGoodListReply_List, 0),
	}

	goods, err, count = uuc.ubRepo.GetGoodsPageThree(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if nil != err {
		return res, nil
	}
	res.Count = count

	for _, v := range goods {
		res.Goods = append(res.Goods, &v1.AdminGoodListReply_List{
			Id:     v.ID,
			Name:   v.Name,
			One:    v.One,
			Two:    v.Two,
			Amount: v.Amount,
			Three:  v.Three,
			Status: v.Status,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*User, error) {
	return uuc.repo.GetUserByUserIds(ctx, userIds...)
}

func (uuc *UserUseCase) GetAllUsers(ctx context.Context) ([]*User, error) {
	return uuc.repo.GetAllUsers(ctx)
}

func (uuc *UserUseCase) AdminUndoUpdate(ctx context.Context, req *v1.AdminUndoUpdateRequest) (*v1.AdminUndoUpdateReply, error) {
	var (
		err  error
		undo int64
	)

	res := &v1.AdminUndoUpdateReply{}

	if 1 == req.SendBody.Undo {
		undo = 1
	} else {
		undo = 0
	}

	_, err = uuc.repo.UndoUser(ctx, req.SendBody.UserId, undo)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) LockUser(ctx context.Context, req *v1.LockUserRequest) (*v1.LockUserReply, error) {
	var (
		err  error
		lock int64
	)

	res := &v1.LockUserReply{}

	if 1 == req.SendBody.Lock {
		lock = 1
	} else {
		lock = 0
	}

	_, err = uuc.repo.LockUser(ctx, req.SendBody.UserId, lock)
	if nil != err {
		return res, err
	}

	if 1 == req.SendBody.One {
		// 推荐
		var (
			userRecommend *UserRecommend
			team          []*UserRecommend
		)
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, req.SendBody.UserId)
		if nil == userRecommend || nil != err {
			return res, nil
		}

		team, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(req.SendBody.UserId, 10))
		if nil != err {
			return res, nil
		}

		for _, v := range team {
			_, err = uuc.repo.LockUser(ctx, v.UserId, lock)
			if nil != err {
				fmt.Println("锁定错误", err, v, lock)
			}
		}
	}

	return res, nil
}

func (uuc *UserUseCase) LockUserReward(ctx context.Context, req *v1.LockUserRewardRequest) (*v1.LockUserRewardReply, error) {
	var (
		err  error
		lock int64
	)

	res := &v1.LockUserRewardReply{}

	if 1 == req.SendBody.LockReward {
		lock = 1
	} else {
		lock = 0
	}

	_, err = uuc.repo.LockUserReward(ctx, req.SendBody.UserId, lock)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminAreaLevelUpdate(ctx context.Context, req *v1.AdminAreaLevelUpdateRequest) (*v1.AdminAreaLevelUpdateReply, error) {
	var (
		err error
	)

	res := &v1.AdminAreaLevelUpdateReply{}

	_, err = uuc.urRepo.UpdateUserAreaLevel(ctx, req.SendBody.UserId, req.SendBody.Level)
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminRecordList(ctx context.Context, req *v1.RecordListRequest) (*v1.RecordListReply, error) {
	var (
		locations  []*EthUserRecord
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.RecordListReply{
		Locations: make([]*v1.RecordListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	locations, err, count = uuc.locationRepo.GetEthUserRecordListByUserId(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}
		tmpCoinType := "RAW"
		res.Locations = append(res.Locations, &v1.RecordListReply_LocationList{
			CreatedAt:  v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:    users[v.UserId].Address,
			AddressTwo: users[v.UserId].AddressTwo,
			Amount:     fmt.Sprintf("%.2f", float64(v.AmountTwo)),
			CoinType:   tmpCoinType,
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	var (
		locationsOut []*Reward
		userSearch   *User
		userId       int64
		userIds      []int64
		userIdsMap   map[int64]int64
		users        map[int64]*User
		count        int64
		err          error
	)

	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	status := "stop"
	if "running" == req.Status {
		var (
			locations []*LocationNew
		)

		status = "running"
		locations, err, count = uuc.locationRepo.GetLocations(ctx, &Pagination{
			PageNum:  int(req.Page),
			PageSize: 10,
		}, userId, status)
		if nil != err {
			return res, nil
		}
		res.Count = count

		userIdsMap = make(map[int64]int64, 0)
		for _, vLocations := range locations {
			userIdsMap[vLocations.UserId] = vLocations.UserId
		}
		for _, v := range userIdsMap {
			userIds = append(userIds, v)
		}

		users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
		if nil != err {
			return res, nil
		}

		for _, v := range locations {
			if _, ok := users[v.UserId]; !ok {
				continue
			}

			var (
				userRecords []*UserBalanceRecord
			)
			userRecords, err = uuc.locationRepo.GetUserBalanceRecordsTwo(ctx, v.UserId)
			if nil != err {
				return res, nil
			}

			var created string
			if 0 < len(userRecords) {
				created = userRecords[0].CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05")
			}
			res.Locations = append(res.Locations, &v1.AdminLocationListReply_LocationList{
				Address:       users[v.UserId].Address,
				Current:       fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
				CurrentMax:    fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
				Out:           users[v.UserId].OutRate,
				CurrentMaxSub: fmt.Sprintf("%.2f", float64(v.CurrentMax-v.Current)/float64(100000)),
				Usdt:          fmt.Sprintf("%.2f", float64(v.Usdt)/float64(100000)),
				CreatedAt:     created,
			})
		}
	} else {
		locationsOut, err, count = uuc.locationRepo.GetLocationsOut(ctx, &Pagination{
			PageNum:  int(req.Page),
			PageSize: 10,
		}, userId, status)
		if nil != err {
			return res, nil
		}
		res.Count = count

		userIdsMap = make(map[int64]int64, 0)
		for _, vLocations := range locationsOut {
			userIdsMap[vLocations.UserId] = vLocations.UserId
		}
		for _, v := range userIdsMap {
			userIds = append(userIds, v)
		}

		users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
		if nil != err {
			return res, nil
		}

		for _, v := range locationsOut {
			if _, ok := users[v.UserId]; !ok {
				continue
			}

			res.Locations = append(res.Locations, &v1.AdminLocationListReply_LocationList{
				Address:       users[v.UserId].Address,
				Current:       fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
				CurrentMax:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(100000)),
				Out:           users[v.UserId].OutRate,
				CurrentMaxSub: "0.00",
				Usdt:          fmt.Sprintf("%.2f", float64(v.Amount)/250000),
				CreatedAt:     v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			})
		}
	}

	return res, nil

}

func (uuc *UserUseCase) AdminLocationListNew(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	var (
		locations  []*LocationNew
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminLocationListReply{
		Locations: make([]*v1.AdminLocationListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	locations, err, count = uuc.locationRepo.GetLocations2(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Locations = append(res.Locations, &v1.AdminLocationListReply_LocationList{
			Address:    users[v.UserId].Address,
			Current:    fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
			CurrentMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminLocationAllList(ctx context.Context, req *v1.AdminLocationAllListRequest) (*v1.AdminLocationAllListReply, error) {
	var (
		locations  []*LocationNew
		userSearch *User
		userId     int64
		userIds    []int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminLocationAllListReply{
		Locations: make([]*v1.AdminLocationAllListReply_LocationList, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	locations, err, count = uuc.locationRepo.GetLocationsAll(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range locations {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range locations {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Locations = append(res.Locations, &v1.AdminLocationAllListReply_LocationList{
			CreatedAt:  v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Address:    users[v.UserId].Address,
			Status:     v.Status,
			Current:    fmt.Sprintf("%.2f", float64(v.Current)/float64(100000)),
			CurrentMax: fmt.Sprintf("%.2f", float64(v.CurrentMax)/float64(100000)),
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminRecommendList(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	var (
		userRecommends []*UserRecommend
		userRecommend  *UserRecommend
		userIdsMap     map[int64]int64
		userIds        []int64
		users          map[int64]*User
		user           *User
		err            error
	)

	res := &v1.AdminUserRecommendReply{
		Users: make([]*v1.AdminUserRecommendReply_List, 0),
	}

	// 地址查询
	if 0 < req.UserId {
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, req.UserId)
		if nil == userRecommend {
			return res, nil
		}

		userRecommends, err = uuc.urRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(userRecommend.UserId, 10))
		if nil != err {
			return res, nil
		}
	} else if "" != req.Address {
		user, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}

		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
		if nil == userRecommend {
			return res, nil
		}

		userRecommends, err = uuc.urRepo.GetUserRecommendByCode(ctx, userRecommend.RecommendCode+"D"+strconv.FormatInt(userRecommend.UserId, 10))
		if nil != err {
			return res, nil
		}
	}

	userIdsMap = make(map[int64]int64, 0)
	for _, vLocations := range userRecommends {
		userIdsMap[vLocations.UserId] = vLocations.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range userRecommends {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Users = append(res.Users, &v1.AdminUserRecommendReply_List{
			Address:   users[v.UserId].Address,
			Id:        v.ID,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.2f", users[v.UserId].MyTotalAmount),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {
	var (
		userCurrentMonthRecommends []*UserCurrentMonthRecommend
		searchUser                 *User
		userIdsMap                 map[int64]int64
		userIds                    []int64
		searchUserId               int64
		users                      map[int64]*User
		count                      int64
		err                        error
	)

	res := &v1.AdminMonthRecommendReply{
		Users: make([]*v1.AdminMonthRecommendReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		searchUser, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil == searchUser {
			return res, nil
		}
		searchUserId = searchUser.ID
	}

	userCurrentMonthRecommends, err, count = uuc.userCurrentMonthRecommendRepo.GetUserCurrentMonthRecommendGroupByUserId(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, searchUserId)
	if nil != err {
		return res, nil
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vRecommend := range userCurrentMonthRecommends {
		userIdsMap[vRecommend.UserId] = vRecommend.UserId
		userIdsMap[vRecommend.RecommendUserId] = vRecommend.RecommendUserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range userCurrentMonthRecommends {
		if _, ok := users[v.UserId]; !ok {
			continue
		}

		res.Users = append(res.Users, &v1.AdminMonthRecommendReply_List{
			Address:          users[v.UserId].Address,
			Id:               v.ID,
			RecommendAddress: users[v.RecommendUserId].Address,
			CreatedAt:        v.Date.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	var (
		configs []*Config
	)

	res := &v1.AdminConfigReply{
		Config: make([]*v1.AdminConfigReply_List, 0),
	}

	configs, _ = uuc.configRepo.GetConfigs(ctx)
	if nil == configs {
		return res, nil
	}

	for _, v := range configs {
		res.Config = append(res.Config, &v1.AdminConfigReply_List{
			Id:    v.ID,
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminConfigUpdateListen(ctx context.Context, req *v1.AdminConfigUpdateListenRequest) (*v1.AdminConfigUpdateListenReply, error) {
	//var (
	//	err         error
	//	priceChange *PriceChange
	//)
	//
	//res := &v1.AdminConfigUpdateListenReply{}
	//priceChange, err = uuc.ubRepo.GetPriceChangeConfig(ctx)
	//if nil != err {
	//	return nil, err
	//}
	//
	//if nil == priceChange {
	//	return res, nil
	//}
	//
	//_, err = uuc.configRepo.UpdatePriceChangeStatus(ctx, priceChange.ID, 1)
	//if nil != err {
	//	return nil, err
	//}
	//
	//var (
	//	configs      []*Config
	//	bPrice       int64
	//	bPriceBase   int64
	//	originBprice int64
	//	//feeRate      int64
	//	users []*User
	//)
	//configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "b_price_base", "exchange_rate")
	//if nil != configs {
	//	for _, vConfig := range configs {
	//		if "b_price_base" == vConfig.KeyName {
	//			bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		}
	//
	//		//if "exchange_rate" == vConfig.KeyName {
	//		//	feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
	//		//}
	//	}
	//}
	//
	//bPrice = priceChange.Price
	//originBprice = priceChange.Origin
	//
	//if 0 >= bPrice || 0 >= bPriceBase {
	//	return nil, err
	//}
	//
	//users, err = uuc.repo.GetAllUsers(ctx)
	//if nil != err {
	//	return nil, err
	//}
	//if nil == users {
	//	return nil, nil
	//}
	//for _, v := range users {
	//	var (
	//		runningLocation *LocationNew
	//		userBalance     *UserBalance
	//	)
	//	runningLocation, err = uuc.locationRepo.GetMyLocationLastRunning(ctx, v.ID)
	//	if nil != err {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	if nil == runningLocation {
	//		continue
	//	}
	//
	//	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, v.ID)
	//	if nil != err {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	if bPrice > originBprice {
	//		// 涨价
	//		tmp := userBalance.BalanceDhb*100/bPriceBase*bPrice - userBalance.BalanceDhb*100/bPriceBase*originBprice
	//		tmp = tmp / 100
	//		if tmp > 0 {
	//
	//			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
	//				runningLocation.Status = "running"
	//				if runningLocation.Current+tmp >= runningLocation.CurrentMax { // 占位分红人分满停止
	//					runningLocation.Status = "stop"
	//					runningLocation.StopDate = time.Now().UTC().Add(8 * time.Hour)
	//
	//					tmp = runningLocation.CurrentMax - runningLocation.Current
	//				}
	//
	//				if 0 < tmp {
	//					var tmpMaxNew int64
	//					if runningLocation.CurrentMax >= runningLocation.CurrentMaxNew {
	//						tmpMaxNew = runningLocation.CurrentMax - runningLocation.CurrentMaxNew
	//					}
	//					err = uuc.locationRepo.UpdateLocationNewNew(ctx, runningLocation.ID, runningLocation.UserId, runningLocation.Status, tmp, tmpMaxNew, 0, runningLocation.StopDate, runningLocation.CurrentMax) // 分红占位数据修改
	//					if nil != err {
	//						return err
	//					}
	//
	//					err = uuc.ubRepo.PriceChange(ctx, runningLocation.UserId, tmp, "up")
	//					if nil != err {
	//						return err
	//					}
	//				}
	//
	//				// 业绩减掉
	//				if "stop" == runningLocation.Status {
	//					//if runningLocation.CurrentMax >= runningLocation.CurrentMaxNew {
	//					//	_, err = uuc.ubRepo.ExchangeBiw(ctx, v.ID, runningLocation.CurrentMax-runningLocation.CurrentMaxNew, feeRate)
	//					//	if nil != err {
	//					//		return err
	//					//	}
	//					//}
	//
	//					tmpTop := runningLocation.Top
	//					tmpTopNum := runningLocation.TopNum
	//					for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
	//						err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, runningLocation.Usdt/100000)
	//						if nil != err {
	//							return err
	//						}
	//
	//						var (
	//							currentLocation *LocationNew
	//						)
	//						currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
	//						if nil != err {
	//							return err
	//						}
	//
	//						if nil != currentLocation && 0 < currentLocation.Top {
	//							tmpTop = currentLocation.Top
	//							tmpTopNum = currentLocation.TopNum
	//							continue
	//						}
	//
	//						break
	//					}
	//				}
	//
	//				return nil
	//			}); nil != err {
	//				fmt.Println("err price change", err, runningLocation)
	//				continue
	//			}
	//		}
	//
	//	} else if bPrice < originBprice {
	//		// 降价
	//		tmp := userBalance.BalanceDhb*100/bPriceBase*originBprice - userBalance.BalanceDhb*100/bPriceBase*bPrice
	//		tmp = tmp / 100
	//		if tmp > 0 {
	//			if runningLocation.Current <= tmp { // 占位分红人分满停止
	//				tmp = runningLocation.Current
	//			}
	//
	//			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
	//				if 0 < tmp {
	//					err = uuc.locationRepo.UpdateLocationNewNewNew(ctx, runningLocation.ID, tmp) // 分红占位数据修改
	//					if nil != err {
	//						return err
	//					}
	//
	//					err = uuc.ubRepo.PriceChange(ctx, runningLocation.UserId, tmp, "down")
	//					if nil != err {
	//						return err
	//					}
	//				}
	//
	//				return nil
	//			}); nil != err {
	//				fmt.Println("err price change", err, runningLocation)
	//				continue
	//			}
	//		}
	//	}
	//}

	return nil, nil
}

func (uuc *UserUseCase) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	var (
		err error
	)

	res := &v1.AdminConfigUpdateReply{}

	var (
		configs []*Config
		bPrice  int64
		//bPriceBase   int64
		originBprice int64
		//feeRate      int64
		//users        []*User
	)
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "b_price", "b_price_base", "exchange_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "b_price" == vConfig.KeyName {
				originBprice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
			//else if "b_price_base" == vConfig.KeyName {
			//	bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//} else if "exchange_rate" == vConfig.KeyName {
			//	feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			//}
		}
	}

	bPrice, _ = strconv.ParseInt(req.SendBody.Value, 10, 64)

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		if 1 == req.SendBody.Id {
			//
			//if 0 >= bPrice || 0 >= bPriceBase {
			//	return nil, err
			//}
			//
			//users, err = uuc.repo.GetAllUsers(ctx)
			//if nil != err {
			//	return nil, err
			//}
			//if nil == users {
			//	return nil, nil
			//}
			//for _, v := range users {
			//	var (
			//		runningLocation *LocationNew
			//		userBalance     *UserBalance
			//	)
			//	runningLocation, err = uuc.locationRepo.GetMyLocationLastRunning(ctx, v.ID)
			//	if nil != err {
			//		fmt.Println(err)
			//		continue
			//	}
			//
			//	if nil == runningLocation {
			//		continue
			//	}
			//
			//	userBalance, err = uuc.ubRepo.GetUserBalance(ctx, v.ID)
			//	if nil != err {
			//		fmt.Println(err)
			//		continue
			//	}
			//
			//	if bPrice > originBprice {
			//		// 涨价
			//		tmp := userBalance.BalanceDhb*100/bPriceBase*bPrice - userBalance.BalanceDhb*100/bPriceBase*originBprice
			//		tmp = tmp / 100
			//		if tmp > 0 {
			//
			//			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//				runningLocation.Status = "running"
			//				if runningLocation.Current+tmp >= runningLocation.CurrentMax { // 占位分红人分满停止
			//					runningLocation.Status = "stop"
			//					runningLocation.StopDate = time.Now().UTC().Add(8 * time.Hour)
			//
			//					tmp = runningLocation.CurrentMax - runningLocation.Current
			//				}
			//
			//				if 0 < tmp {
			//					var tmpMaxNew int64
			//					if runningLocation.CurrentMax >= runningLocation.CurrentMaxNew {
			//						tmpMaxNew = runningLocation.CurrentMax - runningLocation.CurrentMaxNew
			//					}
			//					err = uuc.locationRepo.UpdateLocationNewNew(ctx, runningLocation.ID, runningLocation.Status, tmp, tmpMaxNew, 0, runningLocation.StopDate) // 分红占位数据修改
			//					if nil != err {
			//						return err
			//					}
			//
			//					err = uuc.ubRepo.PriceChange(ctx, runningLocation.UserId, tmp, "up")
			//					if nil != err {
			//						return err
			//					}
			//				}
			//
			//				// 业绩减掉
			//				if "stop" == runningLocation.Status {
			//					if runningLocation.CurrentMax >= runningLocation.CurrentMaxNew {
			//						_, err = uuc.ubRepo.ExchangeBiw(ctx, v.ID, runningLocation.CurrentMax-runningLocation.CurrentMaxNew, feeRate)
			//						if nil != err {
			//							return err
			//						}
			//					}
			//
			//					tmpTop := runningLocation.Top
			//					tmpTopNum := runningLocation.TopNum
			//					for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
			//						err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, runningLocation.Usdt/100000)
			//						if nil != err {
			//							return err
			//						}
			//
			//						var (
			//							currentLocation *LocationNew
			//						)
			//						currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
			//						if nil != err {
			//							return err
			//						}
			//
			//						if nil != currentLocation && 0 < currentLocation.Top {
			//							tmpTop = currentLocation.Top
			//							tmpTopNum = currentLocation.TopNum
			//							continue
			//						}
			//
			//						break
			//					}
			//				}
			//
			//				return nil
			//			}); nil != err {
			//				fmt.Println("err price change", err, runningLocation)
			//				continue
			//			}
			//		}

			//	} else if bPrice < originBprice {
			//		// 降价
			//		tmp := userBalance.BalanceDhb*100/bPriceBase*originBprice - userBalance.BalanceDhb*100/bPriceBase*bPrice
			//		tmp = tmp / 100
			//		if tmp > 0 {
			//			if runningLocation.Current <= tmp { // 占位分红人分满停止
			//				tmp = runningLocation.Current
			//			}
			//
			//			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//				if 0 < tmp {
			//					err = uuc.locationRepo.UpdateLocationNewNewNew(ctx, runningLocation.ID, tmp) // 分红占位数据修改
			//					if nil != err {
			//						return err
			//					}
			//
			//					err = uuc.ubRepo.PriceChange(ctx, runningLocation.UserId, tmp, "down")
			//					if nil != err {
			//						return err
			//					}
			//				}
			//
			//				return nil
			//			}); nil != err {
			//				fmt.Println("err price change", err, runningLocation)
			//				continue
			//			}
			//		}
			//	}
			//}

			err = uuc.configRepo.CreatePriceChangeConfig(ctx, originBprice, bPrice)
			if nil != err {
				return err
			}
		}

		_, err = uuc.configRepo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	//
	//_, err = uuc.configRepo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
	//if nil != err {
	//	return res, err
	//}

	return res, nil
}

func (uuc *UserUseCase) AdminWithdrawPass(ctx context.Context, req *v1.AdminWithdrawPassRequest) (*v1.AdminWithdrawPassReply, error) {
	//var (
	//	err error
	//)

	//res := &v1.AdminWithdrawPassReply{}
	//
	//_, err = uuc.ubRepo.UpdateWithdrawPass(ctx, req.SendBody.Id)
	//if nil != err {
	//	return res, err
	//}

	return nil, nil
}

func (uuc *UserUseCase) AdminPasswordUpdate(ctx context.Context, req *v1.AdminPasswordUpdateRequest) (*v1.AdminPasswordUpdateReply, error) {

	_, _ = uuc.uiRepo.UpdateUserPassword(ctx, req.SendBody.UserId, req.SendBody.Password)
	return &v1.AdminPasswordUpdateReply{}, nil
}

func (uuc *UserUseCase) AdminChangeAddress(ctx context.Context, req *v1.AdminChangeAddressRequest) (*v1.AdminChangeAddressReply, error) {
	var (
		user    *User
		userTwo *User
		err     error
	)
	if req.SendBody.AddressTwo == req.SendBody.Address {
		return nil, err
	}

	user, err = uuc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if nil == user {
		return nil, err
	}

	userTwo, err = uuc.repo.GetUserByAddress(ctx, req.SendBody.AddressTwo)
	if nil != userTwo {
		return nil, err
	}

	return nil, uuc.uiRepo.UpdateUserAddress(ctx, req.SendBody.Address, req.SendBody.AddressTwo)
}

func (uuc *UserUseCase) AdminVipUpdate(ctx context.Context, req *v1.AdminVipUpdateRequest) (*v1.AdminVipUpdateReply, error) {
	var (
		err error
	)

	err = uuc.uiRepo.UpdateUserRecommendLevel2(ctx, req.SendBody.UserId, uint64(req.SendBody.Vip))
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminSetPass(ctx context.Context, req *v1.AdminSetPassRequest) (*v1.AdminSetPassReply, error) {
	var (
		err error
	)

	err = uuc.uiRepo.UpdateUserPass(ctx, req.SendBody.UserId, req.SendBody.Pass)
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminAmountFourUpdate(ctx context.Context, req *v1.AdminAmountFourRequest) (*v1.AdminAmountFourReply, error) {
	var (
		err  error
		user *User
	)

	user, err = uuc.repo.GetUserById(ctx, req.SendBody.UserId)
	if nil != err {
		return nil, err
	}

	var do bool
	if int64(user.AmountFourGet) >= req.SendBody.Amount {
		do = true
	}

	err = uuc.uiRepo.UpdateUserAmountFour(ctx, req.SendBody.UserId, float64(req.SendBody.Amount), do)
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminBalanceUpdate(ctx context.Context, req *v1.AdminBalanceUpdateRequest) (*v1.AdminBalanceUpdateReply, error) {
	var (
		err error
	)
	res := &v1.AdminBalanceUpdateReply{}

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	_, err = uuc.ubRepo.UpdateBalance(ctx, req.SendBody.UserId, amount) // 推荐人信息修改
	if nil != err {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminLogin(ctx context.Context, req *v1.AdminLoginRequest, ca string) (*v1.AdminLoginReply, error) {
	var (
		admin *Admin
		err   error
	)

	res := &v1.AdminLoginReply{}
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	fmt.Println(password)
	admin, err = uuc.repo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil != err {
		return res, err
	}

	claims := auth.CustomClaims{
		UserId:   admin.ID,
		UserType: "admin",
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                     // 签名的生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(48 * time.Hour)), // 2天过期
			Issuer:    "game",
		},
	}

	token, err := auth.CreateToken(claims, ca)
	if err != nil {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "生成token失败")
	}
	res.Token = token
	return res, nil
}

func (uuc *UserUseCase) AdminCreateAccount(ctx context.Context, req *v1.AdminCreateAccountRequest) (*v1.AdminCreateAccountReply, error) {
	var (
		admin    *Admin
		myAdmin  *Admin
		newAdmin *Admin
		err      error
	)

	res := &v1.AdminCreateAccountReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	admin, err = uuc.repo.GetAdminByAccount(ctx, req.SendBody.Account, password)
	if nil != admin {
		return nil, errors.New(500, "ERROR_TOKEN", "已存在账户")
	}

	newAdmin, err = uuc.repo.CreateAdmin(ctx, &Admin{
		Password: password,
		Account:  req.SendBody.Account,
	})

	if nil != newAdmin {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AdminList(ctx context.Context, req *v1.AdminListRequest) (*v1.AdminListReply, error) {
	var (
		admins []*Admin
	)

	res := &v1.AdminListReply{Account: make([]*v1.AdminListReply_List, 0)}

	admins, _ = uuc.repo.GetAdmins(ctx)
	if nil == admins {
		return res, nil
	}

	for _, v := range admins {
		res.Account = append(res.Account, &v1.AdminListReply_List{
			Id:      v.ID,
			Account: v.Account,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AdminChangePassword(ctx context.Context, req *v1.AdminChangePasswordRequest) (*v1.AdminChangePasswordReply, error) {
	var (
		myAdmin *Admin
		admin   *Admin
		err     error
	)

	res := &v1.AdminChangePasswordReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))
	admin, err = uuc.repo.UpdateAdminPassword(ctx, req.SendBody.Account, password)
	if nil == admin {
		return res, err
	}

	return res, nil
}

func (uuc *UserUseCase) AuthList(ctx context.Context, req *v1.AuthListRequest) (*v1.AuthListReply, error) {
	var (
		myAdmin *Admin
		Auths   []*Auth
		err     error
	)

	res := &v1.AuthListReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	Auths, err = uuc.repo.GetAuths(ctx)
	if nil == Auths {
		return res, err
	}

	for _, v := range Auths {
		res.Auth = append(res.Auth, &v1.AuthListReply_List{
			Id:   v.ID,
			Name: v.Name,
			Path: v.Path,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) MyAuthList(ctx context.Context, req *v1.MyAuthListRequest) (*v1.MyAuthListReply, error) {
	var (
		myAdmin   *Admin
		adminAuth []*AdminAuth
		auths     map[int64]*Auth
		authIds   []int64
		err       error
	)

	res := &v1.MyAuthListReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" == myAdmin.Type {
		res.Super = int64(1)
		return res, nil
	}

	adminAuth, err = uuc.repo.GetAdminAuth(ctx, adminId)
	if nil == adminAuth {
		return res, err
	}

	for _, v := range adminAuth {
		authIds = append(authIds, v.AuthId)
	}

	if 0 >= len(authIds) {
		return res, nil
	}

	auths, err = uuc.repo.GetAuthByIds(ctx, authIds...)
	for _, v := range adminAuth {
		if _, ok := auths[v.AuthId]; !ok {
			continue
		}
		res.Auth = append(res.Auth, &v1.MyAuthListReply_List{
			Id:   v.ID,
			Name: auths[v.AuthId].Name,
			Path: auths[v.AuthId].Path,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) UserAuthList(ctx context.Context, req *v1.UserAuthListRequest) (*v1.UserAuthListReply, error) {
	var (
		myAdmin   *Admin
		adminAuth []*AdminAuth
		auths     map[int64]*Auth
		authIds   []int64
		err       error
	)

	res := &v1.UserAuthListReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	adminAuth, err = uuc.repo.GetAdminAuth(ctx, req.AdminId)
	if nil == adminAuth {
		return res, err
	}

	for _, v := range adminAuth {
		authIds = append(authIds, v.AuthId)
	}

	if 0 >= len(authIds) {
		return res, nil
	}

	auths, err = uuc.repo.GetAuthByIds(ctx, authIds...)
	for _, v := range adminAuth {
		if _, ok := auths[v.AuthId]; !ok {
			continue
		}
		res.Auth = append(res.Auth, &v1.UserAuthListReply_List{
			Id:   v.ID,
			Name: auths[v.AuthId].Name,
			Path: auths[v.AuthId].Path,
		})
	}

	return res, nil
}

func (uuc *UserUseCase) AuthAdminCreate(ctx context.Context, req *v1.AuthAdminCreateRequest) (*v1.AuthAdminCreateReply, error) {
	var (
		myAdmin *Admin
		err     error
	)

	res := &v1.AuthAdminCreateReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	_, err = uuc.repo.CreateAdminAuth(ctx, req.SendBody.AdminId, req.SendBody.AuthId)
	if nil != err {
		return nil, errors.New(500, "ERROR_TOKEN", "创建失败")
	}

	return res, err
}

func (uuc *UserUseCase) AuthAdminDelete(ctx context.Context, req *v1.AuthAdminDeleteRequest) (*v1.AuthAdminDeleteReply, error) {
	var (
		myAdmin *Admin
		err     error
	)

	res := &v1.AuthAdminDeleteReply{}

	// 在上下文 context 中取出 claims 对象
	var adminId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		adminId = int64(c["UserId"].(float64))
	}
	myAdmin, err = uuc.repo.GetAdminById(ctx, adminId)
	if nil == myAdmin {
		return res, err
	}
	if "super" != myAdmin.Type {
		return nil, errors.New(500, "ERROR_TOKEN", "非超管")
	}

	_, err = uuc.repo.DeleteAdminAuth(ctx, req.SendBody.AdminId, req.SendBody.AuthId)
	if nil != err {
		return nil, errors.New(500, "ERROR_TOKEN", "删除失败")
	}

	return res, err
}

func (uuc *UserUseCase) GetWithdrawPassOrRewardedFirst(ctx context.Context) (*Withdraw, error) {
	return uuc.ubRepo.GetWithdrawPassOrRewardedFirst(ctx)
}

func (uuc *UserUseCase) GetTradeOk(ctx context.Context) (*Trade, error) {
	return uuc.ubRepo.GetTradeOk(ctx)
}

func (uuc *UserUseCase) UpdateWithdrawDoing(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "doing")
}

func (uuc *UserUseCase) UpdateWithdrawSuccess(ctx context.Context, id int64) (*Withdraw, error) {
	return uuc.ubRepo.UpdateWithdraw(ctx, id, "success")
}

func (uuc *UserUseCase) UpdateTrade(ctx context.Context, id int64) (*Trade, error) {
	return uuc.ubRepo.UpdateTrade(ctx, id, "okk")
}

func (uuc *UserUseCase) UpdateTradeDoing(ctx context.Context, id int64) (*Trade, error) {
	return uuc.ubRepo.UpdateTrade(ctx, id, "doing")
}

func (uuc *UserUseCase) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	var (
		withdraws  []*Withdraw
		userIds    []int64
		userSearch *User
		userId     int64
		userIdsMap map[int64]int64
		users      map[int64]*User
		count      int64
		err        error
	)

	res := &v1.AdminWithdrawListReply{
		Withdraw: make([]*v1.AdminWithdrawListReply_List, 0),
	}

	// 地址查询
	if "" != req.Address {
		userSearch, err = uuc.repo.GetUserByAddress(ctx, req.Address)
		if nil != err {
			return res, nil
		}
		userId = userSearch.ID
	}

	withdraws, err, count = uuc.ubRepo.GetWithdraws(ctx, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, userId, req.WithDrawType)
	if nil != err {
		return res, err
	}
	res.Count = count

	userIdsMap = make(map[int64]int64, 0)
	for _, vWithdraws := range withdraws {
		userIdsMap[vWithdraws.UserId] = vWithdraws.UserId
	}
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}

	users, err = uuc.repo.GetUserByUserIds(ctx, userIds...)
	if nil != err {
		return res, nil
	}

	for _, v := range withdraws {
		if _, ok := users[v.UserId]; !ok {
			continue
		}
		res.Withdraw = append(res.Withdraw, &v1.AdminWithdrawListReply_List{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.Add(8 * time.Hour).Format("2006-01-02 15:04:05"),
			Amount:    fmt.Sprintf("%.4f", v.AmountNew),
			Status:    v.Status,
			Type:      v.Type,
			Address:   users[v.UserId].Address,
			RelAmount: fmt.Sprintf("%.4f", v.RelAmountNew),
		})
	}

	return res, nil

}

func (uuc *UserUseCase) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return &v1.AdminFeeReply{}, nil
}

func (uuc *UserUseCase) AdminFeeDaily(ctx context.Context, req *v1.AdminDailyFeeRequest) (*v1.AdminDailyFeeReply, error) {
	return &v1.AdminDailyFeeReply{}, nil
}

func (uuc *UserUseCase) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {
	var (
		rewards []*Reward
		err     error
		total   *Total
	)
	rewards, err = uuc.ubRepo.GetRewardYes(ctx)
	if nil != err {
		return nil, err
	}

	total, err = uuc.ubRepo.GetTotal(ctx)
	if nil != err {
		return nil, err
	}

	TodayRewardRsdt := float64(0)
	TodayRewardRsdtOther := float64(0)
	TodayWithdraw := float64(0)
	todayDeposit := float64(0)
	for _, v := range rewards {
		if "buy" == v.Reason {
			todayDeposit += v.AmountNew
		}
		if "location" == v.Reason {
			TodayRewardRsdt += v.AmountNew
		}

		if "withdraw" == v.Reason {
			TodayWithdraw += v.AmountNew
		}

		if "recommend" == v.Reason || "recommend_two" == v.Reason || "area" == v.Reason || "area_two" == v.Reason || "all" == v.Reason {
			TodayRewardRsdtOther += v.AmountNew
		}
	}

	var (
		users        []*User
		userBalances []*UserBalance
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil != err {
		return nil, err
	}

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 16 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, -1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}

	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 16, 0, 0, 0, time.UTC)

	totalUserR := int64(0)
	totalUser := int64(0)
	todayUserR := int64(0)
	todayUser := int64(0)
	for _, v := range users {
		totalUserR++
		if 0 < v.Amount || 1 <= v.OutRate {
			totalUser++
		}

		if v.CreatedAt.After(todayStart) && v.CreatedAt.Before(todayEnd) {
			todayUserR++
			if 0 < v.AmountUsdt || 1 <= v.OutRate {
				todayUser++
			}
		}
	}

	balanceUsdtTmp := float64(0)
	balanceIspayTmp := float64(0)
	TotalReward := float64(0)
	userBalances, err = uuc.repo.GetAllUserBalance(ctx)
	if nil != err {
		return nil, err
	}
	for _, v := range userBalances {
		balanceUsdtTmp += v.BalanceUsdtFloat
		balanceIspayTmp += v.BalanceRawFloat

		TotalReward += v.AllFloat + v.RecommendTotalFloat + v.RecommendTotalFloatTwo + v.LocationTotalFloat + v.AreaTotalFloatTwo + v.AreaTotalFloat
	}

	return &v1.AdminAllReply{
		TotalUserR:    totalUserR,
		TotalUser:     totalUser,
		TodayUserR:    todayUserR,
		TodayUser:     todayUser,
		BuyTotal:      fmt.Sprintf("%.2f", total.One),
		TodayBuy:      fmt.Sprintf("%.2f", todayDeposit),
		BalanceUsdt:   fmt.Sprintf("%.2f", balanceUsdtTmp),
		TotalIspay:    fmt.Sprintf("%.2f", balanceIspayTmp),
		TodayOne:      fmt.Sprintf("%.2f", TodayRewardRsdt),
		TodayTwo:      fmt.Sprintf("%.2f", TodayRewardRsdtOther),
		TodayThree:    fmt.Sprintf("%.2f", TodayRewardRsdtOther+TodayRewardRsdtOther),
		TotalReward:   fmt.Sprintf("%.2f", TotalReward),
		TodayWithdraw: fmt.Sprintf("%.2f", TodayWithdraw),
		TotalWithdraw: fmt.Sprintf("%.2f", total.Three),
	}, nil
}

func (uuc *UserUseCase) GetConfigWithdrawDestroyRate(ctx context.Context) ([]*Config, error) {
	return uuc.configRepo.GetConfigByKeys(ctx, "withdraw_destroy_rate")
}

func (uuc *UserUseCase) AdminTrade(ctx context.Context, req *v1.AdminTradeRequest) (*v1.AdminTradeReply, error) {
	//time.Sleep(30 * time.Second) // 错开时间和充值
	var (
		tradeNotDeal                []*Trade
		configs                     []*Config
		withdrawRate                int64
		withdrawRecommendRate       int64
		withdrawRecommendSecondRate int64
		withdrawTeamVipRate         int64
		withdrawTeamVipSecondRate   int64
		withdrawTeamVipThirdRate    int64
		withdrawTeamVipFourthRate   int64
		withdrawTeamVipFifthRate    int64
		withdrawTeamVipLevelRate    int64
		vip0Balance                 int64
		err                         error
	)
	// 配置
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "withdraw_rate",
		"withdraw_recommend_rate", "withdraw_recommend_second_rate",
		"withdraw_team_vip_rate", "withdraw_team_vip_second_rate",
		"withdraw_team_vip_third_rate", "withdraw_team_vip_fourth_rate",
		"withdraw_team_vip_fifth_rate", "withdraw_team_vip_level_rate",
		"withdraw_destroy_rate", "vip_0_balance",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "withdraw_rate" == vConfig.KeyName {
				withdrawRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_recommend_rate" == vConfig.KeyName {
				withdrawRecommendRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_recommend_second_rate" == vConfig.KeyName {
				withdrawRecommendSecondRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_rate" == vConfig.KeyName {
				withdrawTeamVipRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_second_rate" == vConfig.KeyName {
				withdrawTeamVipSecondRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_third_rate" == vConfig.KeyName {
				withdrawTeamVipThirdRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_fourth_rate" == vConfig.KeyName {
				withdrawTeamVipFourthRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_fifth_rate" == vConfig.KeyName {
				withdrawTeamVipFifthRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "withdraw_team_vip_level_rate" == vConfig.KeyName {
				withdrawTeamVipLevelRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_0_balance" == vConfig.KeyName {
				vip0Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	tradeNotDeal, err = uuc.ubRepo.GetTradeNotDeal(ctx)
	if nil == tradeNotDeal {
		return &v1.AdminTradeReply{}, nil
	}

	for _, withdraw := range tradeNotDeal {
		if "default" != withdraw.Status {
			continue
		}

		//if "dhb" == withdraw.Type { // 提现dhb
		//	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		//		_, err = uuc.ubRepo.UpdateWithdrawAmount(ctx, withdraw.ID, "rewarded", currentValue)
		//		if nil != err {
		//			return err
		//		}
		//
		//		return nil
		//	}); nil != err {
		//		return nil, err
		//	}
		//
		//	continue
		//}

		//withdraw.Amount*withdrawRate/100
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			//currentValue -= withdraw.Amount * withdrawRate / 100 // 手续费
			//currentValue -= withdraw.Amount * withdrawDestroyRate / 100
			//fmt.Println(withdraw.Amount, currentValue)
			// 手续费记录
			//err = uuc.ubRepo.SystemFee(ctx, withdraw.Amount*withdrawRate/100, withdraw.ID)
			//if nil != err {
			//	return err
			//}

			_, err = uuc.ubRepo.UpdateTrade(ctx, withdraw.ID, "ok")
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			continue
		}

		var (
			userRecommend       *UserRecommend
			tmpRecommendUserIds []string
		)

		// 推荐人
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, withdraw.UserId)
		if nil == userRecommend {
			continue
		}
		if "" != userRecommend.RecommendCode {
			tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		}

		lastKey := len(tmpRecommendUserIds) - 1
		if 1 > lastKey {
			continue
		}

		lastVip := int64(1)
		level1RewardCount := 1
		level2RewardCount := 1
		level3RewardCount := 1
		level4RewardCount := 1
		level5RewardCount := 1

		withdrawTeamVip := int64(0)
		levelOk := 0
		for i := 0; i <= lastKey; i++ {
			// 有占位信息，推荐人推荐人的上一代
			if lastKey-i <= 0 {
				break
			}

			tmpMyTopUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[lastKey-i], 10, 64) // 最后一位是直推人
			myUserTopRecommendUserInfo, _ := uuc.uiRepo.GetUserInfoByUserId(ctx, tmpMyTopUserRecommendUserId)
			if nil == myUserTopRecommendUserInfo {
				continue
			}
			//
			rewardAmount := withdraw.AmountCsd * withdrawRate / 100
			rewardAmountDhb := withdraw.AmountHbs * withdrawRate / 100
			tmpRecommendUserIdsInt := make([]int64, 0)
			if 1 < lastKey-i {
				for _, va := range tmpRecommendUserIds[1 : lastKey-i] {
					tmpRecommendUserIdsInt1, _ := strconv.ParseInt(va, 10, 64)
					tmpRecommendUserIdsInt = append(tmpRecommendUserIdsInt, tmpRecommendUserIdsInt1)
				}
			}

			if lastVip <= myUserTopRecommendUserInfo.Vip { // 上一个级别比我高
				// 会员团队
				if lastVip < myUserTopRecommendUserInfo.Vip && withdrawTeamVipFifthRate >= withdrawTeamVip {
					var tmp int64
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

						if 2 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipRate

						} else if 3 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipSecondRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipSecondRate

						} else if 4 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipThirdRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipThirdRate

						} else if 5 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipFourthRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipFourthRate

						} else if 6 == myUserTopRecommendUserInfo.Vip {
							tmp = withdrawTeamVipFifthRate - withdrawTeamVip
							withdrawTeamVip = withdrawTeamVipFifthRate
						}

						_, err = uuc.ubRepo.WithdrawNewRewardTeamRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*tmp/100, rewardAmountDhb*tmp/100, withdraw.ID, tmpRecommendUserIdsInt)
						if nil != err {
							return err
						}

						return nil
					}); nil != err {
						continue
					}

					lastVip = myUserTopRecommendUserInfo.Vip
					levelOk = 1
					continue
				}

				// 平级奖
				if 0 < levelOk && lastVip == myUserTopRecommendUserInfo.Vip { // 上一个是vip1和以上且和我平级
					tmpCurrent := 0
					if 2 == myUserTopRecommendUserInfo.Vip {
						if 0 < level1RewardCount {
							tmpCurrent = level1RewardCount
							level1RewardCount--
						}
					} else if 3 == myUserTopRecommendUserInfo.Vip {
						if 0 < level2RewardCount {
							tmpCurrent = level2RewardCount
							level2RewardCount--
						}
					} else if 4 == myUserTopRecommendUserInfo.Vip {
						if 0 < level3RewardCount {
							tmpCurrent = level3RewardCount
							level3RewardCount--
						}
					} else if 5 == myUserTopRecommendUserInfo.Vip {
						if 0 < level4RewardCount {
							tmpCurrent = level4RewardCount
							level4RewardCount--
						}
					} else if 6 == myUserTopRecommendUserInfo.Vip {
						if 0 < level5RewardCount {
							tmpCurrent = level5RewardCount
							level5RewardCount--
						}
					}

					if 0 < tmpCurrent {
						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
							_, err = uuc.ubRepo.WithdrawNewRewardLevelRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*withdrawTeamVipLevelRate/100, rewardAmountDhb*withdrawTeamVipLevelRate/100, withdraw.ID, tmpRecommendUserIdsInt)
							if nil != err {
								return err
							}

							return nil
						}); nil != err {
							continue
						}

						lastVip = myUserTopRecommendUserInfo.Vip
						continue
					}
				}
			}

			if 0 == i { // 当前用户被此人直推

				var userBalance *UserBalance
				userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUserTopRecommendUserInfo.UserId)
				if nil != err {
					continue
				}

				if userBalance.BalanceUsdt/100000 < vip0Balance {
					continue
				}

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					_, err = uuc.ubRepo.WithdrawNewRewardRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*withdrawRecommendRate/100, rewardAmountDhb*withdrawRecommendRate/100, withdraw.ID, tmpRecommendUserIdsInt)
					if nil != err {
						return err
					}

					return nil
				}); nil != err {
					continue
				}

				continue
			} else if 1 == i { // 间接推
				var userBalance *UserBalance
				userBalance, err = uuc.ubRepo.GetUserBalance(ctx, myUserTopRecommendUserInfo.UserId)
				if nil != err {
					continue
				}

				if userBalance.BalanceUsdt/100000 < vip0Balance {
					continue
				}

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					_, err = uuc.ubRepo.WithdrawNewRewardSecondRecommend(ctx, myUserTopRecommendUserInfo.UserId, rewardAmount*withdrawRecommendSecondRate/100, rewardAmountDhb*withdrawRecommendSecondRate/100, withdraw.ID, tmpRecommendUserIdsInt)
					if nil != err {
						return err
					}

					return nil
				}); nil != err {
					continue
				}

				continue
			}
		}
	}

	return &v1.AdminTradeReply{}, nil
}

func (uuc *UserUseCase) AdminDailyBalanceReward(ctx context.Context, req *v1.AdminDailyBalanceRewardRequest) (*v1.AdminDailyBalanceRewardReply, error) {
	var (
		balanceRewards    []*BalanceReward
		configs           []*Config
		balanceRewardRate int64
		coinPrice         int64
		coinRewardRate    int64
		rewardRate        int64
		err               error
	)
	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "balance_reward_rate", "coin_price", "coin_reward_rate", "reward_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "balance_reward_rate" == vConfig.KeyName {
				balanceRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_price" == vConfig.KeyName {
				coinPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_reward_rate" == vConfig.KeyName {
				coinRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "reward_rate" == vConfig.KeyName {
				rewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	now := time.Now()
	if "" != req.Date { // 测试条件
		now, err = time.Parse("2006-01-02 15:04:05", req.Date) // 时间进行格式校验
		if nil != err {
			return nil, err
		}
	}

	now = now.UTC()
	balanceRewards, err = uuc.ubRepo.GetBalanceRewardCurrent(ctx, now)

	timeLimit := time.Now().UTC().Add(-23 * time.Hour)

	for _, vBalanceRewards := range balanceRewards {
		if "" == req.Date { // 测试条件
			if vBalanceRewards.LastRewardDate.After(timeLimit) {
				continue
			}
		}

		// 今天发
		tmpCurrentReward := vBalanceRewards.Amount * balanceRewardRate / 1000
		var myLocationLast *LocationNew
		// 获取当前用户的占位信息，已经有运行中的跳过
		myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vBalanceRewards.UserId)
		if nil == myLocationLast { // 无占位信息
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpCurrentStatus := myLocationLast.Status // 现在还在运行中

			tmpBalanceUsdtAmount := tmpCurrentReward * rewardRate / 100 // 记录下一次
			tmpBalanceCoinAmount := tmpCurrentReward * coinRewardRate / 100 * 1000 / coinPrice

			myLocationLast.Status = "running"
			myLocationLast.Current += tmpCurrentReward
			if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
				if "running" == tmpCurrentStatus {
					myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

					lastRewardAmount := tmpCurrentReward - (myLocationLast.Current - myLocationLast.CurrentMax)
					tmpBalanceUsdtAmount = lastRewardAmount * rewardRate / 100 // 记录下一次
					tmpBalanceCoinAmount = lastRewardAmount * coinRewardRate / 100 * 1000 / coinPrice
				}
				myLocationLast.Status = "stop"
			}

			if 0 < tmpCurrentReward {
				err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, tmpCurrentReward, myLocationLast.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}

				_, err = uuc.ubRepo.UserDailyBalanceReward(ctx, vBalanceRewards.UserId, tmpCurrentReward, tmpBalanceUsdtAmount, tmpBalanceCoinAmount, tmpCurrentStatus)
				if nil != err {
					return err
				}

				err = uuc.ubRepo.UpdateBalanceRewardLastRewardDate(ctx, vBalanceRewards.ID)
				if nil != err {
					return err
				}
			}

			return nil
		}); nil != err {
			continue
		}

	}

	return &v1.AdminDailyBalanceRewardReply{}, nil
}

func (uuc *UserUseCase) AdminDailyLocationRewardNewTwo(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		level1   float64
		level2   float64
		level3   float64
		level4   float64
		level5   float64
		vip1     float64
		vip2     float64
		vip3     float64
		vip4     float64
		vip5     float64
		vip6     float64
		vip7     float64
		vip8     float64
		vip9     float64
		vip10    float64
		b1       float64
		fourRate float64
		configs  []*Config
		err      error
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "level_2", "level_3", "level_4", "level_5", "level_1", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v9", "v8", "v10", "b_1", "four_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "level_1" == vConfig.KeyName {
				level1, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_2" == vConfig.KeyName {
				level2, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_3" == vConfig.KeyName {
				level3, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_4" == vConfig.KeyName {
				level4, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_5" == vConfig.KeyName {
				level5, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v1" == vConfig.KeyName {
				vip1, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v2" == vConfig.KeyName {
				vip2, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v3" == vConfig.KeyName {
				vip3, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v4" == vConfig.KeyName {
				vip4, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v5" == vConfig.KeyName {
				vip5, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v6" == vConfig.KeyName {
				vip6, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v7" == vConfig.KeyName {
				vip7, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v8" == vConfig.KeyName {
				vip8, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v9" == vConfig.KeyName {
				vip9, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v10" == vConfig.KeyName {
				vip10, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "b_1" == vConfig.KeyName {
				b1, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "four_rate" == vConfig.KeyName {
				fourRate, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	var (
		users       []*User
		usersMap    map[int64]*User
		stopUserIds map[int64]bool
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		fmt.Println("今日分红错误用户获取失败")
		return nil, nil
	}

	stopUserIds = make(map[int64]bool, 0)
	usersMap = make(map[int64]*User, 0)

	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
		myLowUser         map[int64][]*UserRecommend
	)

	myLowUser = make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)

	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		buyRecords []*BuyRecord
	)

	//buyRecords, err = uuc.repo.GetBuyRecord(ctx, -9)
	buyRecords, err = uuc.repo.GetBuyRecord(ctx, -9)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败")
		return nil, nil
	}

	// 静态
	for _, vBuyRecords := range buyRecords {
		if _, ok := usersMap[vBuyRecords.UserId]; !ok {
			continue
		}

		tmpUsers := usersMap[vBuyRecords.UserId]

		// 出局的
		if 0 >= tmpUsers.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[tmpUsers.ID]; ok {
			continue
		}

		num := float64(0)
		numTwo := float64(0)
		if 1 == tmpUsers.Last {
			num = level1
			numTwo = 1.5
		} else if 2 == tmpUsers.Last {
			num = level2
			numTwo = 2
		} else if 3 == tmpUsers.Last {
			num = level3
			numTwo = 2.5
		} else if 4 == tmpUsers.Last {
			num = level4
			numTwo = 3
		} else if 5 == tmpUsers.Last {
			num = level5
			numTwo = 3.5
		} else {
			continue
		}

		stop := false
		tmp := vBuyRecords.Amount * num
		if tmp+tmpUsers.AmountUsdtGet >= tmpUsers.AmountUsdt*numTwo {
			tmp = math.Abs(tmpUsers.AmountUsdt*numTwo - tmpUsers.AmountUsdtGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
			userRecommend = userRecommendsMap[tmpUsers.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, tmpUsers)
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserRewardNew(ctx, vBuyRecords.ID, tmpUsers.ID, tmp, tmpUsers.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, tmpUsers)
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, tmpUsers)
			continue
		}

		if stop {
			stopUserIds[tmpUsers.ID] = true // 出局

			if nil != userRecommend && "" != userRecommend.RecommendCode {
				var tmpRecommendUserIds []string
				tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
				for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
					if 0 >= len(tmpRecommendUserIds[j]) {
						continue
					}

					myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
					if 0 >= myUserRecommendUserId {
						continue
					}
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, tmpUsers.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, tmpUsers, myUserRecommendUserId)
							return err
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily 业绩更新", err, tmpUsers)
						continue
					}

					// 级别降低
					// 我的下级，更新vip
					userIdsLowTmpTwo := make([]int64, 0)
					for _, vTmpLow := range myLowUser[myUserRecommendUserId] {
						userIdsLowTmpTwo = append(userIdsLowTmpTwo, vTmpLow.UserId)
					}
					if 0 < len(userIdsLowTmpTwo) {
						uuc.updateVip(ctx, myUserRecommendUserId, userIdsLowTmpTwo)
					}
				}
			}

		}
	}

	tmpMapCurrentI := make(map[int]float64, 0)
	tmpB1 := b1
	for i := 1; i <= 18; i++ {
		if i > 1 {
			tmpMapCurrentI[i] = tmpB1 / 2
		} else {
			tmpMapCurrentI[i] = tmpB1
		}
		tmpB1 = tmpMapCurrentI[i]
	}

	// 帮扶
	for _, vBuyRecords := range buyRecords {
		if 0 < vBuyRecords.AmountGet {
			continue
		}

		if _, ok := usersMap[vBuyRecords.UserId]; !ok {
			continue
		}

		if 1 == usersMap[vBuyRecords.UserId].LockReward {
			continue
		}

		tmpUsers := usersMap[vBuyRecords.UserId]

		// 出局的
		if 0 >= tmpUsers.AmountUsdt {
			continue
		}

		//num := float64(0)
		////numTwo := float64(0)
		//if 1 == tmpUsers.Last {
		//	num = level1
		//	//numTwo = 1.5
		//} else if 2 == tmpUsers.Last {
		//	num = level2
		//	//numTwo = 1.8
		//} else if 3 == tmpUsers.Last {
		//	num = level3
		//	//numTwo = 2
		//} else if 4 == tmpUsers.Last {
		//	num = level4
		//	//numTwo = 2.5
		//} else if 5 == tmpUsers.Last {
		//	num = level5
		//	//numTwo = 3
		//} else {
		//	continue
		//}

		tmp := vBuyRecords.Amount
		//if tmp+tmpUsers.AmountUsdtGet >= tmpUsers.AmountUsdt*numTwo {
		//	tmp = math.Abs(tmpUsers.AmountUsdt*numTwo - tmpUsers.AmountUsdtGet)
		//}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
			userRecommend = userRecommendsMap[tmpUsers.ID]
		} else {
			fmt.Println("错误分红帮扶，信息缺失：", err, tmpUsers)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		currentI := 1
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			if currentI > 18 {
				break
			}

			tmpCurrentI := currentI
			currentI++

			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红帮扶，信息缺失,user：", err, tmpUsers)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红帮扶，信息缺失,user1：", err, tmpUsers)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红帮扶，信息缺失3：", err, tmpUserId, tmpUsers)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红帮扶，信息缺失3：", err, tmpUserId, tmpUsers)
				continue
			}

			// 条件
			if tmpCurrentI < 10 {
				if tmpCurrentI > len(myLowUser[tmpUserId]) {
					continue
				}
			} else if 10 > len(myLowUser[tmpUserId]) {
				continue
			}

			if _, ok := tmpMapCurrentI[tmpCurrentI]; !ok {
				fmt.Println("错误分红帮扶，信息缺失35：", err, tmpUserId, tmpUsers)
				continue
			}
			tmpRecommendAmount := tmp * tmpMapCurrentI[tmpCurrentI]

			var (
				stopRecommend   bool
				numRecommendTwo float64
			)
			if 1 == tmpRecommendUser.Last {
				numRecommendTwo = 1.5
			} else if 2 == tmpRecommendUser.Last {
				numRecommendTwo = 2
			} else if 3 == tmpRecommendUser.Last {
				numRecommendTwo = 2.5
			} else if 4 == tmpRecommendUser.Last {
				numRecommendTwo = 3
			} else if 5 == tmpRecommendUser.Last {
				numRecommendTwo = 3.5
			} else {
				continue
			}

			if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*numRecommendTwo {
				tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*numRecommendTwo - tmpRecommendUser.AmountUsdtGet)
				stopRecommend = true
			}

			// 分红
			tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
			if 0 >= tmpRecommendAmount {
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardRecommendNew(ctx, tmpRecommendUser.ID, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, int64(tmpCurrentI), tmpUsers.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily recommend 18", err, tmpRecommendUser)
			}

			if stopRecommend {
				stopUserIds[tmpRecommendUser.ID] = true // 出局

				// 推荐人
				var (
					userRecommendArea *UserRecommend
				)
				if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
					userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
				} else {
					fmt.Println("错误分红帮扶，信息缺失7：", err, tmpRecommendUser)
				}

				if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					var tmpRecommendAreaUserIds []string
					tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

					for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendAreaUserIds[j]) {
							continue
						}

						myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendAreaUserId {
							continue
						}

						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily 业绩更新", err, tmpUsers)
							continue
						}

						// 级别降低
						// 我的下级，更新vip
						userIdsLowTmpTwo := make([]int64, 0)
						for _, vTmpLow := range myLowUser[myUserRecommendAreaUserId] {
							userIdsLowTmpTwo = append(userIdsLowTmpTwo, vTmpLow.UserId)
						}
						if 0 < len(userIdsLowTmpTwo) {
							uuc.updateVip(ctx, myUserRecommendAreaUserId, userIdsLowTmpTwo)
						}
					}
				}
			}
		}
	}

	// 小区
	for _, vBuyRecords := range buyRecords {
		if _, ok := usersMap[vBuyRecords.UserId]; !ok {
			continue
		}

		if 1 == usersMap[vBuyRecords.UserId].LockReward {
			continue
		}

		tmpUsers := usersMap[vBuyRecords.UserId]

		// 出局的
		if 0 >= tmpUsers.AmountUsdt {
			continue
		}

		num := float64(0)
		//numTwo := float64(0)
		if 1 == tmpUsers.Last {
			num = level1
			//numTwo = 1.5
		} else if 2 == tmpUsers.Last {
			num = level2
			//numTwo = 1.8
		} else if 3 == tmpUsers.Last {
			num = level3
			//numTwo = 2
		} else if 4 == tmpUsers.Last {
			num = level4
			//numTwo = 2.5
		} else if 5 == tmpUsers.Last {
			num = level5
			//numTwo = 3
		} else {
			continue
		}

		tmp := vBuyRecords.Amount * num
		//if tmp+tmpUsers.AmountUsdtGet >= tmpUsers.AmountUsdt*numTwo {
		//	tmp = math.Abs(tmpUsers.AmountUsdt*numTwo - tmpUsers.AmountUsdtGet)
		//}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpUsers.ID]; ok {
			userRecommend = userRecommendsMap[tmpUsers.ID]
		} else {
			fmt.Println("错误分红帮扶，信息缺失：", err, tmpUsers)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		currentI := 1
		lastLevel := int64(0)
		lastNumArea := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			tmpCurrentI := currentI
			currentI++

			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, tmpUsers)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, tmpUsers)
				continue
			}

			tmpVip := int64(0)
			if 0 < tmpRecommendUser.VipAdmin {
				tmpVip = tmpRecommendUser.VipAdmin
			} else {
				tmpVip = tmpRecommendUser.Vip
			}

			tmpNumArea := float64(0)
			if 1 == tmpVip {
				tmpNumArea = vip1
			} else if 2 == tmpVip {
				tmpNumArea = vip2
			} else if 3 == tmpVip {
				tmpNumArea = vip3
			} else if 4 == tmpVip {
				tmpNumArea = vip4
			} else if 5 == tmpVip {
				tmpNumArea = vip5
			} else if 6 == tmpVip {
				tmpNumArea = vip6
			} else if 7 == tmpVip {
				tmpNumArea = vip7
			} else if 8 == tmpVip {
				tmpNumArea = vip8
			} else if 9 == tmpVip {
				tmpNumArea = vip9
			} else if 10 == tmpVip {
				tmpNumArea = vip10
			} else {
				continue
			}

			if 0 >= tmpNumArea {
				continue
			}

			if lastLevel >= tmpVip {
				continue
			}

			// 出局沉淀
			if 0 >= tmpRecommendUser.AmountUsdt {
				lastLevel = tmpVip
				lastNumArea = tmpNumArea
				continue
			}

			if tmpNumArea <= lastNumArea {
				continue
			}

			tmpCurrentNum := tmpNumArea - lastNumArea

			lastLevel = tmpVip
			lastNumArea = tmpNumArea

			tmpRecommendAmount := tmp * tmpCurrentNum

			var (
				stopRecommend   bool
				numRecommendTwo float64
			)
			if 1 == tmpRecommendUser.Last {
				numRecommendTwo = 1.5
			} else if 2 == tmpRecommendUser.Last {
				numRecommendTwo = 2
			} else if 3 == tmpRecommendUser.Last {
				numRecommendTwo = 2.5
			} else if 4 == tmpRecommendUser.Last {
				numRecommendTwo = 3
			} else if 5 == tmpRecommendUser.Last {
				numRecommendTwo = 3.5
			} else {
				continue
			}

			if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*numRecommendTwo {
				tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*numRecommendTwo - tmpRecommendUser.AmountUsdtGet)
				stopRecommend = true
			}

			// 分红
			tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
			if 0 >= tmpRecommendAmount {
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardAreaNew(ctx, tmpRecommendUser.ID, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, tmpVip, int64(tmpCurrentI), tmpUsers.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily recommend 18", err, tmpRecommendUser)
			}

			if stopRecommend {
				stopUserIds[tmpRecommendUser.ID] = true // 出局

				// 推荐人
				var (
					userRecommendArea *UserRecommend
				)
				if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
					userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
				} else {
					fmt.Println("错误分红帮扶，信息缺失7：", err, tmpRecommendUser)
				}

				if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					var tmpRecommendAreaUserIds []string
					tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

					for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendAreaUserIds[j]) {
							continue
						}

						myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendAreaUserId {
							continue
						}

						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily 业绩更新", err, tmpUsers)
							continue
						}

						// 级别降低
						// 我的下级，更新vip
						userIdsLowTmpTwo := make([]int64, 0)
						for _, vTmpLow := range myLowUser[myUserRecommendAreaUserId] {
							userIdsLowTmpTwo = append(userIdsLowTmpTwo, vTmpLow.UserId)
						}
						if 0 < len(userIdsLowTmpTwo) {
							uuc.updateVip(ctx, myUserRecommendAreaUserId, userIdsLowTmpTwo)
						}
					}
				}
			}
		}
	}

	for _, vUsers := range users {
		if 0 >= vUsers.AmountFour {
			continue
		}

		if vUsers.AmountFour <= vUsers.AmountFourGet {
			continue
		}

		tmp := fourRate * vUsers.AmountFour
		if tmp+vUsers.AmountFourGet >= vUsers.AmountFour {
			tmp = math.Abs(vUsers.AmountFour - vUsers.AmountFourGet)
		}

		tmp = math.Round(tmp*10000000) / 10000000
		if 0 >= tmp {
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserRewardNewFour(ctx, vUsers.ID, tmp)
			if code > 0 && err != nil {
				fmt.Println("错误基金：", err, vUsers)
			}

			return nil
		}); nil != err {
			fmt.Println("err reward jj daily", err, vUsers)
			continue
		}
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminDailyLocationRewardNewThree(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		allOne  float64
		allTwo  float64
		configs []*Config
		err     error
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "all_one", "all_two")
	if nil != configs {
		for _, vConfig := range configs {
			if "all_one" == vConfig.KeyName {
				allOne, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "all_two" == vConfig.KeyName {
				allTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
		}
	}

	var (
		users       []*User
		usersMap    map[int64]*User
		stopUserIds map[int64]bool
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		fmt.Println("今日分红错误用户获取失败")
		return nil, nil
	}

	stopUserIds = make(map[int64]bool, 0)
	usersMap = make(map[int64]*User, 0)

	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
		myLowUser         map[int64][]*UserRecommend
	)

	myLowUser = make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)

	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		total *Total
	)
	total, err = uuc.ubRepo.GetTotal(ctx)
	if nil == total {
		fmt.Println("今日分红错误用户获取失败，total")
		return nil, nil
	}

	if 0 >= total.One {
		fmt.Println("今日分红0，total", total)
		return nil, nil
	}

	var (
		usersOrderAmountBiw []*User
	)
	usersOrderAmountBiw, err = uuc.repo.GetAllUsersOrderAmountBiw(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败，total，推荐人数")
		return nil, nil
	}

	var (
		usersOrderRecommendOrder []*User
	)
	usersOrderRecommendOrder, err = uuc.repo.GetAllUsersRecommendOrder(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败，total，推荐1人数")
		return nil, nil
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserRewardTotalOver(ctx)
		if err != nil {
			fmt.Println("错误分红帮扶：", err)
		}

		return nil
	}); nil != err {
		fmt.Println("err reward daily recommend over total", err)
	}

	tmpReward := total.One * allOne
	if 0 < tmpReward {
		// 全球
		for k, v := range usersOrderAmountBiw {
			if 0 >= v.AmountUsdt {
				continue
			}

			tmpRecommendAmount := float64(0)
			if 0 == k {
				tmpRecommendAmount = tmpReward * 0.5
			} else if 1 == k {
				tmpRecommendAmount = tmpReward * 0.3
			} else if 2 == k {
				tmpRecommendAmount = tmpReward * 0.2
			} else {
				break
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[v.ID]; ok {
				continue
			}

			tmpRecommendUser := v
			if nil == tmpRecommendUser {
				fmt.Println("错误分红全球1，信息缺失,user1：", err, tmpRecommendUser)
				continue
			}

			var (
				stopRecommend   bool
				numRecommendTwo float64
			)
			if 1 == tmpRecommendUser.Last {
				numRecommendTwo = 1.5
			} else if 2 == tmpRecommendUser.Last {
				numRecommendTwo = 2
			} else if 3 == tmpRecommendUser.Last {
				numRecommendTwo = 2.5
			} else if 4 == tmpRecommendUser.Last {
				numRecommendTwo = 3
			} else if 5 == tmpRecommendUser.Last {
				numRecommendTwo = 3.5
			} else {
				continue
			}

			if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*numRecommendTwo {
				tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*numRecommendTwo - tmpRecommendUser.AmountUsdtGet)
				stopRecommend = true
			}

			// 分红
			tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
			if 0 >= tmpRecommendAmount {
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardTotalOneNew(ctx, tmpRecommendUser.ID, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, int64(k))
				if code > 0 && err != nil {
					fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily recommend 18", err, tmpRecommendUser)
			}

			if stopRecommend {
				stopUserIds[tmpRecommendUser.ID] = true // 出局

				// 推荐人
				var (
					userRecommendArea *UserRecommend
				)
				if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
					userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
				} else {
					fmt.Println("错误分红帮扶，信息缺失7：", err, tmpRecommendUser)
				}

				if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					var tmpRecommendAreaUserIds []string
					tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

					for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendAreaUserIds[j]) {
							continue
						}

						myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendAreaUserId {
							continue
						}

						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily 业绩更新", err, tmpRecommendUser)
							continue
						}

						// 级别降低
						// 我的下级，更新vip
						userIdsLowTmpTwo := make([]int64, 0)
						for _, vTmpLow := range myLowUser[myUserRecommendAreaUserId] {
							userIdsLowTmpTwo = append(userIdsLowTmpTwo, vTmpLow.UserId)
						}
						if 0 < len(userIdsLowTmpTwo) {
							uuc.updateVip(ctx, myUserRecommendAreaUserId, userIdsLowTmpTwo)
						}
					}
				}
			}
		}
	} else {
		fmt.Println("无全球分红基金", total, allOne)
	}

	tmpRewardTwo := total.One * allTwo
	if 0 < tmpRewardTwo {
		// 全球
		for k, v := range usersOrderRecommendOrder {
			if 0 >= v.AmountUsdt {
				continue
			}

			tmpRecommendAmount := float64(0)
			if 0 == k {
				tmpRecommendAmount = tmpRewardTwo * 0.5
			} else if 1 == k {
				tmpRecommendAmount = tmpRewardTwo * 0.3
			} else if 2 == k {
				tmpRecommendAmount = tmpRewardTwo * 0.2
			} else {
				break
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[v.ID]; ok {
				continue
			}

			tmpRecommendUser := v
			if nil == tmpRecommendUser {
				fmt.Println("错误分红全球1，信息缺失,user1：", err, tmpRecommendUser)
				continue
			}

			var (
				stopRecommend   bool
				numRecommendTwo float64
			)
			if 1 == tmpRecommendUser.Last {
				numRecommendTwo = 1.5
			} else if 2 == tmpRecommendUser.Last {
				numRecommendTwo = 2
			} else if 3 == tmpRecommendUser.Last {
				numRecommendTwo = 2.5
			} else if 4 == tmpRecommendUser.Last {
				numRecommendTwo = 3
			} else if 5 == tmpRecommendUser.Last {
				numRecommendTwo = 3.5
			} else {
				continue
			}

			if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*numRecommendTwo {
				tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*numRecommendTwo - tmpRecommendUser.AmountUsdtGet)
				stopRecommend = true
			}

			// 分红
			tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
			if 0 >= tmpRecommendAmount {
				continue
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardTotalTwoNew(ctx, tmpRecommendUser.ID, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, int64(k))
				if code > 0 && err != nil {
					fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily recommend 18", err, tmpRecommendUser)
			}

			if stopRecommend {
				stopUserIds[tmpRecommendUser.ID] = true // 出局

				// 推荐人
				var (
					userRecommendArea *UserRecommend
				)
				if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
					userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
				} else {
					fmt.Println("错误分红帮扶，信息缺失7：", err, tmpRecommendUser)
				}

				if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					var tmpRecommendAreaUserIds []string
					tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

					for j := len(tmpRecommendAreaUserIds) - 1; j >= 0; j-- {
						if 0 >= len(tmpRecommendAreaUserIds[j]) {
							continue
						}

						myUserRecommendAreaUserId, _ := strconv.ParseInt(tmpRecommendAreaUserIds[j], 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendAreaUserId {
							continue
						}

						if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红帮扶：", err, tmpRecommendUser)
							}

							return nil
						}); nil != err {
							fmt.Println("err reward daily 业绩更新", err, tmpRecommendUser)
							continue
						}

						// 级别降低
						// 我的下级，更新vip
						userIdsLowTmpTwo := make([]int64, 0)
						for _, vTmpLow := range myLowUser[myUserRecommendAreaUserId] {
							userIdsLowTmpTwo = append(userIdsLowTmpTwo, vTmpLow.UserId)
						}
						if 0 < len(userIdsLowTmpTwo) {
							uuc.updateVip(ctx, myUserRecommendAreaUserId, userIdsLowTmpTwo)
						}
					}
				}
			}
		}
	} else {
		fmt.Println("无全球分红基金", total, allTwo)
	}

	return nil, nil
}

func (uuc *UserUseCase) updateVip(ctx context.Context, tmpUserId int64, userIdsLowTmp []int64) {
	var (
		err error
	)

	// 下级的级别
	var (
		tmpLowUsers map[int64]*User
	)
	tmpLowUsers, err = uuc.repo.GetUserByUserIdsTwo(ctx, userIdsLowTmp)
	if err != nil {
		fmt.Println("update vip 遍历业绩2：", err)
		return
	}

	vip9 := 0
	vip8 := 0
	vip7 := 0
	vip6 := 0
	vip5 := 0
	vip4 := 0
	vip3 := 0
	vip2 := 0
	vip1 := 0
	// 获取业绩
	tmpAreaMax := float64(0)
	tmpAreaMin := float64(0)
	tmpMaxId := int64(0)
	for _, vMyLowUser := range tmpLowUsers {
		if 9 == vMyLowUser.Vip {
			vip9++
		} else if 8 == vMyLowUser.Vip {
			vip8++
		} else if 7 == vMyLowUser.Vip {
			vip7++
		} else if 6 == vMyLowUser.Vip {
			vip6++
		} else if 5 == vMyLowUser.Vip {
			vip5++
		} else if 4 == vMyLowUser.Vip {
			vip4++
		} else if 3 == vMyLowUser.Vip {
			vip3++
		} else if 2 == vMyLowUser.Vip {
			vip2++
		} else if 1 == vMyLowUser.Vip {
			vip1++
		}

		if tmpAreaMax < vMyLowUser.MyTotalAmount+float64(vMyLowUser.AmountSelf) {
			tmpAreaMax = vMyLowUser.MyTotalAmount + float64(vMyLowUser.AmountSelf)
			tmpMaxId = vMyLowUser.ID
		}
	}

	if 0 < tmpMaxId {
		for _, vMyLowUser := range tmpLowUsers {
			if tmpMaxId != vMyLowUser.ID {
				tmpAreaMin += vMyLowUser.MyTotalAmount + float64(vMyLowUser.AmountSelf)
			}
		}
	}

	tmpVip := int64(0)
	if 10000000 <= tmpAreaMin && 2 <= vip9 {
		tmpVip = 10
	} else if 6000000 <= tmpAreaMin && 2 <= vip8+vip9 {
		tmpVip = 9
	} else if 3000000 <= tmpAreaMin && 2 <= vip7+vip8+vip9 {
		tmpVip = 8
	} else if 1300000 <= tmpAreaMin && 2 <= vip6+vip7+vip8+vip9 {
		tmpVip = 7
	} else if 600000 <= tmpAreaMin && 2 <= vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 6
	} else if 300000 <= tmpAreaMin && 2 <= vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 5
	} else if 100000 <= tmpAreaMin && 2 <= vip3+vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 4
	} else if 50000 <= tmpAreaMin && 2 <= vip2+vip3+vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 3
	} else if 20000 <= tmpAreaMin && 2 <= vip1+vip2+vip3+vip4+vip5+vip6+vip7+vip8+vip9 {
		tmpVip = 2
	} else if 5000 <= tmpAreaMin {
		tmpVip = 1
	}

	// 增加业绩
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.repo.UpdateUserVip(ctx, tmpUserId, tmpVip)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println("update vip 遍历业绩 vip：", err)
		return
	}

	return
}

func (uuc *UserUseCase) AdminMyTotalAmount(ctx context.Context, req *v1.AdminDailyRewardRequest) (*v1.AdminDailyRewardReply, error) {

	var (
		allRecord []*BuyRecord
		err       error
	)
	allRecord, err = uuc.uiRepo.GetAllBuyRecord(ctx)
	if nil != err {
		return nil, nil
	}

	//fmt.Println("总计", len(allRecord), "条")
	//var (
	//	users    []*User
	//	usersMap map[int64]*User
	//)
	//users, err = uuc.ubRepo.GetAllUsersB(ctx)
	//if nil == users {
	//	return nil, nil
	//}
	//
	//usersMap = make(map[int64]*User, 0)
	//for _, vUsers := range users {
	//	usersMap[vUsers.ID] = vUsers
	//}

	//var (
	//	userRecommends    []*UserRecommend
	//	userRecommendsMap map[int64]*UserRecommend
	//)
	//userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	//if nil != err {
	//	return nil, nil
	//}
	//
	//userRecommendsMap = make(map[int64]*UserRecommend, 0)
	//for _, vUr := range userRecommends {
	//	userRecommendsMap[vUr.UserId] = vUr
	//}
	//
	//userTotalAmount := make(map[int64]float64, 0)
	//tmpAmount := float64(0)
	for _, v := range allRecord {

		err = uuc.uiRepo.UpdateUserAmountSelf(ctx, v.UserId, uint64(v.Amount))
		if nil != err {
			fmt.Println(err)
		}
		//tmpAmount += v.Amount
		//if _, ok := userRecommendsMap[v.UserId]; !ok {
		//	continue
		//}
		//
		//tmpRecommendUserIds := make([]string, 0)
		//userRecommend := userRecommendsMap[v.UserId]
		//if "" != userRecommend.RecommendCode {
		//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		//}
		//
		//totalTmp := len(tmpRecommendUserIds) - 1
		//for i := totalTmp; i >= 0; i-- {
		//	tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
		//	if 0 >= tmpUserId {
		//		continue
		//	}
		//
		//	userTotalAmount[tmpUserId] += v.Amount
		//}

	}

	//fmt.Println("总计", tmpAmount, "usdt")
	//
	//for k, v := range userTotalAmount {
	//	fmt.Println("用户", k, v)
	//}

	return nil, nil
}

func (uuc *UserUseCase) AdminDailyRewardFour(ctx context.Context, req *v1.AdminDailyRewardRequest) (*v1.AdminDailyRewardReply, error) {
	lockAll.Lock()
	defer lockAll.Unlock()
	var (
		configs    []*Config
		openReward uint64
		err        error
	)
	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx, "open_reward")
	if nil != err || nil == configs {
		fmt.Println("错误分红，配置", err)
		return &v1.AdminDailyRewardReply{}, nil
	}
	for _, vConfig := range configs {
		if "open_reward" == vConfig.KeyName {
			openReward, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	if 1 != openReward {
		fmt.Println("关闭分红", openReward)
		return &v1.AdminDailyRewardReply{}, nil
	}

	var (
		buyRecords []*BuyRecordFour
	)
	buyRecords, err = uuc.uiRepo.GetBuyRecordFour(ctx)
	if nil != err {
		fmt.Println("认购数据查询错误")
		return nil, err
	}

	for _, v := range buyRecords {
		stop := false
		tmpC := float64(0)
		if v.Amount <= v.AmountGet {
			fmt.Println("错误数据，每日ispay分红", v)
			continue
		}

		if v.Amount <= v.AmountGetPerDay+v.AmountGet {
			stop = true
			tmpC = v.Amount - v.AmountGet
		} else {
			tmpC = v.AmountGetPerDay
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.uiRepo.UpdateUserRewardDailyIspay(ctx, v.ID, v.UserId, tmpC, v.Amount, float64(v.Four), stop)
			if err != nil {
				fmt.Println("错误分红静态：", err, v)
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminDailyReward(ctx context.Context, req *v1.AdminDailyRewardRequest) (*v1.AdminDailyRewardReply, error) {
	lockAll.Lock()
	defer lockAll.Unlock()

	var (
		configs             []*Config
		oneRate             float64
		twoRate             float64
		threeRate           float64
		fourRate            float64
		fiveRate            float64
		sixRate             float64
		sevenRate           float64
		eightRate           float64
		nineRate            float64
		oneOneRate          float64
		oneTwoRate          float64
		recommendTwoRate    float64
		recommendTwoRateSub float64
		areaOne             float64
		areaTwo             float64
		areaThree           float64
		areaFour            float64
		areaFive            float64
		areaZero            float64
		allEach             float64
		uRate               float64
		bRate               float64
		openReward          uint64
		err                 error
	)
	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"buy_one", "buy_two", "buy_three", "buy_four", "buy_five", "buy_six", "buy_seven", "buy_eight", "buy_nine",
		"buy_one_one", "buy_one_two",
		"recommend_two", "recommend_two_sub",
		"area_one", "area_two", "area_three", "area_four", "area_five", "area_zero",
		"all_each",
		"u_rate",
		"b_rate",
		"open_reward",
	)
	if nil != err || nil == configs {
		fmt.Println("错误分红，配置", err)
		return &v1.AdminDailyRewardReply{}, nil
	}
	for _, vConfig := range configs {
		if "buy_one" == vConfig.KeyName {
			oneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_two" == vConfig.KeyName {
			twoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_three" == vConfig.KeyName {
			threeRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_four" == vConfig.KeyName {
			fourRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_five" == vConfig.KeyName {
			fiveRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_six" == vConfig.KeyName {
			sixRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_seven" == vConfig.KeyName {
			sevenRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_eight" == vConfig.KeyName {
			eightRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_nine" == vConfig.KeyName {
			nineRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_one_one" == vConfig.KeyName {
			oneOneRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "buy_one_two" == vConfig.KeyName {
			oneTwoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "recommend_two" == vConfig.KeyName {
			recommendTwoRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "recommend_two_sub" == vConfig.KeyName {
			recommendTwoRateSub, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_one" == vConfig.KeyName {
			areaOne, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_two" == vConfig.KeyName {
			areaTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_three" == vConfig.KeyName {
			areaThree, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_four" == vConfig.KeyName {
			areaFour, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_five" == vConfig.KeyName {
			areaFive, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "area_zero" == vConfig.KeyName {
			areaZero, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "all_each" == vConfig.KeyName {
			allEach, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "u_rate" == vConfig.KeyName {
			uRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "b_rate" == vConfig.KeyName {
			bRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "open_reward" == vConfig.KeyName {
			openReward, _ = strconv.ParseUint(vConfig.Value, 10, 64)
		}
	}

	if 1 != openReward {
		fmt.Println("关闭分红", openReward)
		return &v1.AdminDailyRewardReply{}, nil
	}

	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
	)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		return &v1.AdminDailyRewardReply{}, nil
	}

	myLowUser := make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)
	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	var (
		users       []*User
		usersReward []*User
		usersMap    map[int64]*User
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		return &v1.AdminDailyRewardReply{}, nil
	}

	usersMap = make(map[int64]*User, 0)
	usersReward = make([]*User, 0)
	levelOne := make([]*User, 0)
	levelTwo := make([]*User, 0)
	levelThree := make([]*User, 0)
	levelFour := make([]*User, 0)
	levelFive := make([]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers

		if 0 < vUsers.Amount {
			usersReward = append(usersReward, vUsers)
		}
	}

	for _, vUsers := range users {
		if 1 == vUsers.Lock {
			continue
		}

		if 0 < vUsers.VipAdmin && vUsers.Amount > 0 {
			if 1 == vUsers.VipAdmin {
				levelOne = append(levelOne, vUsers)
			} else if 2 == vUsers.VipAdmin {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
			} else if 3 == vUsers.VipAdmin {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
				levelThree = append(levelThree, vUsers)
			} else if 4 == vUsers.VipAdmin {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
				levelThree = append(levelThree, vUsers)
				levelFour = append(levelFour, vUsers)
			} else if 5 == vUsers.VipAdmin {
				levelOne = append(levelOne, vUsers)
				levelTwo = append(levelTwo, vUsers)
				levelThree = append(levelThree, vUsers)
				levelFour = append(levelFour, vUsers)
				levelFive = append(levelFive, vUsers)
			} else {
				// 跳过，没级别
				continue
			}

			continue
		}

		if 1 >= len(myLowUser[vUsers.ID]) {
			continue
		}

		// 获取业绩
		tmpAreaMax := uint64(0)
		tmpAreaMin := uint64(0)
		tmpMaxId := int64(0)
		for _, vMyLowUser := range myLowUser[vUsers.ID] {
			if _, ok := usersMap[vMyLowUser.UserId]; !ok {
				continue
			}

			if tmpAreaMax < uint64(usersMap[vMyLowUser.UserId].MyTotalAmount)+usersMap[vMyLowUser.UserId].AmountSelf {
				tmpAreaMax = uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].AmountSelf
				tmpMaxId = vMyLowUser.ID
			}
		}

		if 0 >= tmpMaxId {
			continue
		}

		for _, vMyLowUser := range myLowUser[vUsers.ID] {
			if _, ok := usersMap[vMyLowUser.UserId]; !ok {
				continue
			}

			if tmpMaxId != vMyLowUser.ID {
				tmpAreaMin += uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].AmountSelf
			}
		}

		if 1500000 <= tmpAreaMin && vUsers.Amount > 0 {
			levelFive = append(levelFive, vUsers)
			levelFour = append(levelFour, vUsers)
			levelThree = append(levelThree, vUsers)
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 500000 <= tmpAreaMin && vUsers.Amount > 0 {
			levelFour = append(levelFour, vUsers)
			levelThree = append(levelThree, vUsers)
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 150000 <= tmpAreaMin && vUsers.Amount > 0 {
			levelThree = append(levelThree, vUsers)
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 50000 <= tmpAreaMin && vUsers.Amount > 0 {
			levelTwo = append(levelTwo, vUsers)
			levelOne = append(levelOne, vUsers)
		} else if 10000 <= tmpAreaMin && vUsers.Amount > 0 {
			levelOne = append(levelOne, vUsers)
		} else {
			// 跳过，没级别
			continue
		}
	}

	stopUserIds := make(map[int64]bool, 0)

	var (
		buyRecords []*BuyRecord
	)
	buyRecords, err = uuc.uiRepo.GetBuyRecord(ctx, 0)
	if nil != err {
		fmt.Println("认购数据查询错误")
		return nil, err
	}

	userBuyRecords := make(map[int64][]*BuyRecord, 0)
	for _, v := range buyRecords {
		if _, ok := userBuyRecords[v.UserId]; !ok {
			userBuyRecords[v.UserId] = make([]*BuyRecord, 0)
		}

		userBuyRecords[v.UserId] = append(userBuyRecords[v.UserId], v)
	}

	t := time.Date(2026, 2, 18, 14, 0, 0, 0, time.UTC)
	// 静态
	for _, tmpBuyRecords := range buyRecords {
		if _, ok := usersMap[tmpBuyRecords.UserId]; !ok {
			continue
		}
		if 1 == usersMap[tmpBuyRecords.UserId].Lock {
			continue
		}

		num := 2.5
		if tmpBuyRecords.CreatedAt.After(t) {
			amountB := uint64(tmpBuyRecords.Amount)
			if 4999 <= amountB && 15001 > amountB {
				num = 3
			} else if 29999 <= amountB && 50001 > amountB {
				num = 3.5
			} else if 99999 <= amountB && 150001 > amountB {
				num = 4
			}
		}
		if tmpBuyRecords.Amount*num <= tmpBuyRecords.AmountGet {
			fmt.Println("错误的数据，已经最大却没停，daily", tmpBuyRecords)
			continue
		}

		numTwo := float64(0)
		if 150000 <= tmpBuyRecords.Amount {
			numTwo = oneTwoRate
		} else if 100000 <= tmpBuyRecords.Amount {
			numTwo = oneOneRate
		} else if 50000 <= tmpBuyRecords.Amount {
			numTwo = nineRate
		} else if 30000 <= tmpBuyRecords.Amount {
			numTwo = eightRate
		} else if 15000 <= tmpBuyRecords.Amount {
			numTwo = sevenRate
		} else if 10000 <= tmpBuyRecords.Amount {
			numTwo = sixRate
		} else if 5000 <= tmpBuyRecords.Amount {
			numTwo = fiveRate
		} else if 1000 <= tmpBuyRecords.Amount {
			numTwo = fourRate
		} else if 500 <= tmpBuyRecords.Amount {
			numTwo = threeRate
		} else if 300 <= tmpBuyRecords.Amount {
			numTwo = twoRate
		} else if 50 <= tmpBuyRecords.Amount {
			numTwo = oneRate
		} else {
			continue
		}

		stop := false
		tmp := tmpBuyRecords.Amount * numTwo
		if tmp+tmpBuyRecords.AmountGet >= tmpBuyRecords.Amount*num {
			tmp = math.Abs(tmpBuyRecords.Amount*num - tmpBuyRecords.AmountGet)
			tmpBuyRecords.AmountGet = tmpBuyRecords.Amount * num
			stop = true
		} else {
			tmpBuyRecords.AmountGet += tmp
		}

		tmpURel := math.Round(tmp*uRate*10000000) / 10000000
		tmpB := math.Round(tmp*bRate*10000000) / 10000000

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			err = uuc.uiRepo.UpdateUserRewardDailyLocation(ctx, tmpBuyRecords.ID, tmpBuyRecords.UserId, tmpURel, tmpB, tmp, tmpBuyRecords.Amount, stop)
			if err != nil {
				fmt.Println("错误分红静态：", err, tmpBuyRecords)
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, tmpBuyRecords)
			continue
		}

		if stop {
			stopUserIds[tmpBuyRecords.ID] = true // 出局

			//// 推荐人
			//var (
			//	userRecommend *UserRecommend
			//)
			//if _, ok := userRecommendsMap[tmpBuyRecords.UserId]; ok {
			//	userRecommend = userRecommendsMap[tmpBuyRecords.UserId]
			//} else {
			//	fmt.Println("错误分红静态，信息缺失：", err, v)
			//	continue
			//}
			//
			//if nil != userRecommend && "" != userRecommend.RecommendCode {
			//	var tmpRecommendUserIds []string
			//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
			//	for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
			//		if 0 >= len(tmpRecommendUserIds[j]) {
			//			continue
			//		}
			//
			//		myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
			//		if 0 >= myUserRecommendUserId {
			//			continue
			//		}
			//		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
			//			// 减掉业绩
			//			err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, tmpBuyRecords.Amount)
			//			if err != nil {
			//				fmt.Println("错误分红静态：", err, tmpBuyRecords, myUserRecommendUserId)
			//				return err
			//			}
			//
			//			return nil
			//		}); nil != err {
			//			fmt.Println("err reward daily 业绩更新", err, tmpBuyRecords)
			//			continue
			//		}
			//	}
			//}

		}
	}

	// 团队和平级
	for _, tmpBuyRecords := range buyRecords {
		if _, ok := usersMap[tmpBuyRecords.UserId]; !ok {
			continue
		}
		if 1 == usersMap[tmpBuyRecords.UserId].Lock {
			continue
		}

		num := 2.5
		if _, ok := usersMap[tmpBuyRecords.UserId]; !ok {
			continue
		}

		if 1 == usersMap[tmpBuyRecords.UserId].LockReward {
			continue
		}

		numTwo := float64(0)
		if 150000 <= tmpBuyRecords.Amount {
			numTwo = oneTwoRate
		} else if 100000 <= tmpBuyRecords.Amount {
			numTwo = oneOneRate
		} else if 50000 <= tmpBuyRecords.Amount {
			numTwo = nineRate
		} else if 30000 <= tmpBuyRecords.Amount {
			numTwo = eightRate
		} else if 15000 <= tmpBuyRecords.Amount {
			numTwo = sevenRate
		} else if 10000 <= tmpBuyRecords.Amount {
			numTwo = sixRate
		} else if 5000 <= tmpBuyRecords.Amount {
			numTwo = fiveRate
		} else if 1000 <= tmpBuyRecords.Amount {
			numTwo = fourRate
		} else if 500 <= tmpBuyRecords.Amount {
			numTwo = threeRate
		} else if 300 <= tmpBuyRecords.Amount {
			numTwo = twoRate
		} else if 50 <= tmpBuyRecords.Amount {
			numTwo = oneRate
		} else {
			continue
		}

		tmp := tmpBuyRecords.Amount * numTwo
		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpBuyRecords.UserId]; ok {
			userRecommend = userRecommendsMap[tmpBuyRecords.UserId]
		} else {
			fmt.Println("错误分红团队，信息缺失：", err, tmpBuyRecords)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		lastKey := len(tmpRecommendUserIds) - 1
		tmpI := int64(0)
		tmpLevelSame := uint64(0)
		for i := lastKey; i >= 0; i-- {
			currentLevel := 0
			tmpI++

			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, tmpBuyRecords)
				continue
			}

			if 1 == usersMap[tmpUserId].Lock {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, tmpBuyRecords)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, tmpBuyRecords)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, tmpBuyRecords)
				continue
			}

			if 0 >= tmpRecommendUser.Amount {
				continue
			}

			if _, ok := userBuyRecords[tmpUserId]; !ok {
				continue
			}

			if 0 >= len(userBuyRecords[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := uint64(0)
			tmpAreaMin := uint64(0)
			tmpMaxId := int64(0)
			tmpMaxUserId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, tmpBuyRecords)
					continue
				}

				if tmpAreaMax < uint64(usersMap[vMyLowUser.UserId].MyTotalAmount)+usersMap[vMyLowUser.UserId].AmountSelf {
					tmpAreaMax = uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].AmountSelf
					tmpMaxId = vMyLowUser.ID
					tmpMaxUserId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId || 0 >= tmpMaxUserId {
				continue
			}

			//fmt.Println("测试1", tmpBuyRecords.ID, tmpUserId, tmpAreaMax, tmpMaxUserId)
			// 如果是我大区的人，不拿，当前人的下级是不是大区的用户id
			if i == lastKey {
				// 直推，是我的大区
				if tmpMaxUserId == tmpBuyRecords.UserId {
					//fmt.Println("测试1：", tmpUserId, tmpMaxId, v.ID)
					continue
				}

				//fmt.Println("测试2：", tmpUserId, tmpMaxId, v.ID)
			} else {
				if i+1 > lastKey {
					fmt.Println("错误分红小区，信息缺失44：", err, tmpUserId, lastKey, i+1, tmpBuyRecords)
					continue
				}

				tmpLastUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i+1], 10, 64) // 最后一位是直推人
				if 0 >= tmpLastUserId {
					fmt.Println("错误分红小区，信息缺失445：", err, tmpUserId, lastKey, i+1, tmpBuyRecords)
					continue
				}

				// 是我大区的人，跳过
				if tmpMaxUserId == tmpLastUserId {
					//fmt.Println("测试3：", tmpUserId, tmpMaxId, tmpLastUserId)
					continue
				}

				//fmt.Println("测试4：", tmpUserId, tmpMaxId, tmpLastUserId)
			}

			//fmt.Println("测试2", tmpBuyRecords.ID, tmpUserId, tmpAreaMax, tmpMaxUserId)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, tmpBuyRecords)
					continue
				}

				if tmpMaxId != vMyLowUser.ID {
					tmpAreaMin += uint64(usersMap[vMyLowUser.UserId].MyTotalAmount) + usersMap[vMyLowUser.UserId].AmountSelf
				}
			}

			//fmt.Println("测试3", tmpBuyRecords.ID, tmpUserId, tmpAreaMax, tmpMaxUserId, tmpAreaMin)
			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.VipAdmin {
				if 1 == tmpRecommendUser.VipAdmin {
					currentLevel = 1
					tmpLastLevelNum = areaOne
				} else if 2 == tmpRecommendUser.VipAdmin {
					currentLevel = 2
					tmpLastLevelNum = areaTwo
				} else if 3 == tmpRecommendUser.VipAdmin {
					currentLevel = 3
					tmpLastLevelNum = areaThree
				} else if 4 == tmpRecommendUser.VipAdmin {
					currentLevel = 4
					tmpLastLevelNum = areaFour
				} else if 5 == tmpRecommendUser.VipAdmin {
					currentLevel = 5
					tmpLastLevelNum = areaFive
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1500000 <= tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = areaFive
				} else if 500000 <= tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = areaFour
				} else if 150000 <= tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = areaThree
				} else if 50000 <= tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = areaTwo
				} else if 10000 <= tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = areaOne
				} else {
					// 跳过，没级别
					continue
				}
			}

			var (
				tmpAreaAmount float64
			)
			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmpAreaAmount = tmp * areaZero
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, tmpBuyRecords, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmpAreaAmount = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			if 0 >= tmpAreaAmount {
				continue
			}

			// 平级，结束
			tmpLevel := false
			if currentLevel == lastLevel {
				tmpLevel = true
				if 5 == tmpLevelSame {
					fmt.Println("5次平级", tmpRecommendUser)
					continue
				}
				tmpLevelSame++
			}

			for _, vUserRecords := range userBuyRecords[tmpUserId] {
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}

				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停，recommend", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)
				tmpU := tmpAreaAmount
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				tmpAreaAmount -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardAreaOne(ctx, vUserRecords.ID, tmpUserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea, usersMap[tmpBuyRecords.UserId].Address, tmpI, int64(currentLevel), tmpLevel)
					if err != nil {
						fmt.Println("错误分红小区：", err, tmpRecommendUser)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily area", err, tmpBuyRecords)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
					//	userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					//} else {
					//	fmt.Println("错误分红小区，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红小区：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < tmpAreaAmount {
						continue
					}
				}

				break
			}

			// 平级，结束
			//if tmpLevel {
			//	break
			//}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}

		}

	}

	// 直推加速
	for _, tmpBuyRecords := range buyRecords {
		num := 2.5
		if _, ok := usersMap[tmpBuyRecords.UserId]; !ok {
			continue
		}

		if 1 == usersMap[tmpBuyRecords.UserId].Lock {
			continue
		}

		if 1 == usersMap[tmpBuyRecords.UserId].LockReward {
			continue
		}

		numTwo := float64(0)
		if 150000 <= tmpBuyRecords.Amount {
			numTwo = oneTwoRate
		} else if 100000 <= tmpBuyRecords.Amount {
			numTwo = oneOneRate
		} else if 50000 <= tmpBuyRecords.Amount {
			numTwo = nineRate
		} else if 30000 <= tmpBuyRecords.Amount {
			numTwo = eightRate
		} else if 15000 <= tmpBuyRecords.Amount {
			numTwo = sevenRate
		} else if 10000 <= tmpBuyRecords.Amount {
			numTwo = sixRate
		} else if 5000 <= tmpBuyRecords.Amount {
			numTwo = fiveRate
		} else if 1000 <= tmpBuyRecords.Amount {
			numTwo = fourRate
		} else if 500 <= tmpBuyRecords.Amount {
			numTwo = threeRate
		} else if 300 <= tmpBuyRecords.Amount {
			numTwo = twoRate
		} else if 50 <= tmpBuyRecords.Amount {
			numTwo = oneRate
		} else {
			continue
		}

		tmp := tmpBuyRecords.Amount * numTwo
		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[tmpBuyRecords.UserId]; ok {
			userRecommend = userRecommendsMap[tmpBuyRecords.UserId]
		} else {
			fmt.Println("错误分红团队，信息缺失：", err, tmpBuyRecords)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		tmpI := 0
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			if 10 == tmpI {
				break
			}

			tmpI++

			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红加速，信息缺失,user：", err, tmpBuyRecords)
				continue
			}

			if 1 == usersMap[tmpUserId].Lock {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红加速，信息缺失3：", err, tmpUserId, tmpBuyRecords)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红加速，信息缺失3：", err, tmpUserId, tmpBuyRecords)
				continue
			}

			// 2代1个，依次类推
			tmpLen := 0
			for _, vTmp := range myLowUser[tmpUserId] {
				if 0 >= usersMap[vTmp.UserId].Amount && 0 >= usersMap[vTmp.UserId].OutRate {
					continue
				}

				tmpLen++
			}

			if tmpLen < tmpI-1 {
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红加速，信息缺失,user1：", err, tmpBuyRecords)
				continue
			}

			if 0 >= tmpRecommendUser.Amount {
				continue
			}

			var (
				tmpAreaAmount float64
				tmpRate       = 0.01
			)

			if recommendTwoRate >= recommendTwoRateSub*float64(tmpI-1) {
				tmpRate = recommendTwoRate - recommendTwoRateSub*float64(tmpI-1)
			}

			tmpAreaAmount = tmp * tmpRate
			if 0 >= tmpAreaAmount {
				continue
			}

			if _, ok := userBuyRecords[tmpUserId]; !ok {
				continue
			}

			for _, vUserRecords := range userBuyRecords[tmpUserId] {
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}

				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停，加速", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)
				tmpU := tmpAreaAmount
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				tmpAreaAmount -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardRecommendNewTwo(ctx, vUserRecords.ID, tmpUserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea, usersMap[tmpBuyRecords.UserId].Address, int64(tmpI))
					if err != nil {
						fmt.Println("错误分红加速：", err, tmpRecommendUser)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily speed", err, tmpBuyRecords)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
					//	userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					//} else {
					//	fmt.Println("错误分红speed，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红speed：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < tmpAreaAmount {
						continue
					}
				}

				break
			}
		}

	}

	var (
		reward []*Reward
	)
	reward, err = uuc.repo.GetRewardYes(ctx)
	if nil != err {
		return &v1.AdminDailyRewardReply{}, nil
	}

	tmpAmountAll := float64(0)
	for _, vReward := range reward {
		tmpAmountAll += vReward.AmountNew
	}

	// 昨日入金: 500100 11 10 10 6 6
	fmt.Println("昨日入金:", tmpAmountAll, len(levelOne), len(levelTwo), len(levelThree), len(levelFour), len(levelFive))
	if 0 >= tmpAmountAll {
		return &v1.AdminDailyRewardReply{}, nil
	}

	tmpRewardAllEach := tmpAmountAll * allEach
	if 0 >= tmpRewardAllEach {
		return &v1.AdminDailyRewardReply{}, nil
	}

	if 0 < len(levelOne) {
		tmpLevelC := tmpRewardAllEach / 5 / float64(len(levelOne))

		for _, v := range levelOne {
			tmpUsers := v

			if _, ok := userBuyRecords[tmpUsers.ID]; !ok {
				continue
			}

			levelTmp := tmpLevelC
			for _, vUserRecords := range userBuyRecords[tmpUsers.ID] {
				num := 2.5
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}
				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停， all1", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)

				tmpU := levelTmp
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				levelTmp -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardAllNew(ctx, vUserRecords.ID, vUserRecords.UserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea)
					if err != nil {
						fmt.Println("错误分红all：", err, vUserRecords)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily all", err, v)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[vUserRecords.UserId]; ok {
					//	userRecommendArea = userRecommendsMap[vUserRecords.UserId]
					//} else {
					//	fmt.Println("错误分红all，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红all：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < levelTmp {
						continue
					}
				}

				break
			}
		}
	}

	if 0 < len(levelTwo) {
		tmpLevelC := tmpRewardAllEach / 5 / float64(len(levelTwo))

		for _, v := range levelTwo {
			tmpUsers := v

			if _, ok := userBuyRecords[tmpUsers.ID]; !ok {
				continue
			}

			levelTmp := tmpLevelC
			for _, vUserRecords := range userBuyRecords[tmpUsers.ID] {
				num := 2.5
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}
				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停，all2", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)

				tmpU := levelTmp
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				levelTmp -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardAllNew(ctx, vUserRecords.ID, vUserRecords.UserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea)
					if err != nil {
						fmt.Println("错误分红all：", err, vUserRecords)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily all", err, v)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[vUserRecords.UserId]; ok {
					//	userRecommendArea = userRecommendsMap[vUserRecords.UserId]
					//} else {
					//	fmt.Println("错误分红all，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红all：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < levelTmp {
						continue
					}
				}

				break
			}
		}
	}

	if 0 < len(levelThree) {
		tmpLevelC := tmpRewardAllEach / 5 / float64(len(levelThree))

		for _, v := range levelThree {
			tmpUsers := v

			if _, ok := userBuyRecords[tmpUsers.ID]; !ok {
				continue
			}

			levelTmp := tmpLevelC
			for _, vUserRecords := range userBuyRecords[tmpUsers.ID] {
				num := 2.5
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}
				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停，all3", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)

				tmpU := levelTmp
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				levelTmp -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardAllNew(ctx, vUserRecords.ID, vUserRecords.UserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea)
					if err != nil {
						fmt.Println("错误分红all：", err, vUserRecords)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily all", err, v)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[vUserRecords.UserId]; ok {
					//	userRecommendArea = userRecommendsMap[vUserRecords.UserId]
					//} else {
					//	fmt.Println("错误分红all，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红all：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < levelTmp {
						continue
					}
				}

				break
			}
		}
	}

	if 0 < len(levelFour) {
		tmpLevelC := tmpRewardAllEach / 5 / float64(len(levelFour))

		for _, v := range levelFour {
			tmpUsers := v

			if _, ok := userBuyRecords[tmpUsers.ID]; !ok {
				continue
			}

			levelTmp := tmpLevelC
			for _, vUserRecords := range userBuyRecords[tmpUsers.ID] {
				num := 2.5
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}
				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停，all4", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)

				tmpU := levelTmp
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				levelTmp -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardAllNew(ctx, vUserRecords.ID, vUserRecords.UserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea)
					if err != nil {
						fmt.Println("错误分红all：", err, vUserRecords)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily all", err, v)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[vUserRecords.UserId]; ok {
					//	userRecommendArea = userRecommendsMap[vUserRecords.UserId]
					//} else {
					//	fmt.Println("错误分红all，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红all：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < levelTmp {
						continue
					}
				}

				break
			}
		}
	}

	if 0 < len(levelFive) {
		tmpLevelC := tmpRewardAllEach / 5 / float64(len(levelFive))

		for _, v := range levelFive {
			tmpUsers := v

			if _, ok := userBuyRecords[tmpUsers.ID]; !ok {
				continue
			}

			levelTmp := tmpLevelC
			for _, vUserRecords := range userBuyRecords[tmpUsers.ID] {
				num := 2.5
				// 本次执行已经出局
				if _, ok := stopUserIds[vUserRecords.ID]; ok {
					continue
				}

				if vUserRecords.CreatedAt.After(t) {
					amountB := uint64(vUserRecords.Amount)
					if 4999 <= amountB && 15001 > amountB {
						num = 3
					} else if 29999 <= amountB && 50001 > amountB {
						num = 3.5
					} else if 99999 <= amountB && 150001 > amountB {
						num = 4
					}
				}
				if vUserRecords.Amount*num <= vUserRecords.AmountGet {
					fmt.Println("错误的数据，已经最大却没停，all5", vUserRecords)
					continue
				}

				var (
					stopArea bool
				)

				tmpU := levelTmp
				if tmpU+vUserRecords.AmountGet >= vUserRecords.Amount*num {
					tmpU = math.Abs(vUserRecords.Amount*num - vUserRecords.AmountGet)
					vUserRecords.AmountGet = vUserRecords.Amount * num
					stopArea = true
				} else {
					vUserRecords.AmountGet += tmpU
				}

				levelTmp -= tmpU

				tmpURel := math.Round(tmpU*uRate*10000000) / 10000000
				tmpB := math.Round(tmpU*bRate*10000000) / 10000000

				if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

					err = uuc.uiRepo.UpdateUserRewardAllNew(ctx, vUserRecords.ID, vUserRecords.UserId, tmpURel, tmpB, tmpU, vUserRecords.Amount, stopArea)
					if err != nil {
						fmt.Println("错误分红all：", err, vUserRecords)
						return err
					}

					return nil
				}); nil != err {
					fmt.Println("err reward daily all", err, v)
				}

				if stopArea {
					stopUserIds[vUserRecords.ID] = true // 出局

					//// 推荐人
					//var (
					//	userRecommendArea *UserRecommend
					//)
					//if _, ok := userRecommendsMap[vUserRecords.UserId]; ok {
					//	userRecommendArea = userRecommendsMap[vUserRecords.UserId]
					//} else {
					//	fmt.Println("错误分红all，信息缺失7：", err, v)
					//}
					//
					//if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
					//	var tmpRecommendAreaUserIds []string
					//	tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
					//
					//	for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
					//		if 0 >= len(vTmpRecommendAreaUserIds) {
					//			continue
					//		}
					//
					//		myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
					//		if 0 >= myUserRecommendAreaUserId {
					//			continue
					//		}
					//
					//		// 减掉业绩
					//		err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendAreaUserId, vUserRecords.Amount)
					//		if err != nil {
					//			fmt.Println("错误分红all：", err, v)
					//		}
					//	}
					//}

					if 0.000001 < levelTmp {
						continue
					}
				}

				break
			}
		}
	}

	return &v1.AdminDailyRewardReply{}, nil
}

func (uuc *UserUseCase) AdminDailyLocationReward(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		level1 float64
		level2 float64
		level3 float64
		level4 float64
		level5 float64
		level6 float64
		vv1    float64
		v2     float64
		v3     float64
		v4     float64
		v5     float64
		v6     float64
		v7     float64
		v8     float64
		v0     float64
		//va4     float64
		//va5     float64
		//va6     float64
		//va7     float64
		//va8     float64
		configs []*Config
		err     error
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "level_2", "level_3", "level_4", "level_6", "level_5", "level_1", "v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8", "v0", "va4", "va5", "va6", "va7", "va8")
	if nil != configs {
		for _, vConfig := range configs {
			if "level_1" == vConfig.KeyName {
				level1, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_2" == vConfig.KeyName {
				level2, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_3" == vConfig.KeyName {
				level3, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_4" == vConfig.KeyName {
				level4, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_5" == vConfig.KeyName {
				level5, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "level_6" == vConfig.KeyName {
				level6, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v1" == vConfig.KeyName {
				vv1, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v2" == vConfig.KeyName {
				v2, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v3" == vConfig.KeyName {
				v3, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v4" == vConfig.KeyName {
				v4, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v5" == vConfig.KeyName {
				v5, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v6" == vConfig.KeyName {
				v6, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v7" == vConfig.KeyName {
				v7, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v8" == vConfig.KeyName {
				v8, _ = strconv.ParseFloat(vConfig.Value, 10)
			} else if "v0" == vConfig.KeyName {
				v0, _ = strconv.ParseFloat(vConfig.Value, 10)
			}
			//else if "va4" == vConfig.KeyName {
			//	va4, _ = strconv.ParseFloat(vConfig.Value, 10)
			//} elseif "va5" == vConfig.KeyName {
			//	va5, _ = strconv.ParseFloat(vConfig.Value, 10)
			//} else if "va6" == vConfig.KeyName {
			//	va6, _ = strconv.ParseFloat(vConfig.Value, 10)
			//} else if "va7" == vConfig.KeyName {
			//	va7, _ = strconv.ParseFloat(vConfig.Value, 10)
			//} else if "va8" == vConfig.KeyName {
			//	va8, _ = strconv.ParseFloat(vConfig.Value, 10)
			//}
		}
	}

	var (
		users       []*User
		usersMap    map[int64]*User
		stopUserIds map[int64]bool
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		fmt.Println("今日分红错误用户获取失败")
		return nil, nil
	}

	stopUserIds = make(map[int64]bool, 0)
	usersMap = make(map[int64]*User, 0)

	userReward1 := make([]*User, 0)
	userReward2 := make([]*User, 0)
	userReward3 := make([]*User, 0)
	userReward4 := make([]*User, 0)
	userReward5 := make([]*User, 0)
	userReward6 := make([]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers

		// 出局的
		if 0 >= vUsers.AmountUsdt {
			continue
		}

		if 1 == vUsers.Last {
			userReward1 = append(userReward1, vUsers)
		} else if 2 == vUsers.Last {
			userReward2 = append(userReward2, vUsers)
		} else if 3 == vUsers.Last {
			userReward3 = append(userReward3, vUsers)
		} else if 4 == vUsers.Last {
			userReward4 = append(userReward4, vUsers)
		} else if 5 == vUsers.Last {
			userReward5 = append(userReward5, vUsers)
		} else if 6 == vUsers.Last {
			userReward6 = append(userReward6, vUsers)
		} else {
			continue
		}
	}

	// 推荐人
	var (
		userRecommends    []*UserRecommend
		userRecommendsMap map[int64]*UserRecommend
		myLowUser         map[int64][]*UserRecommend
	)

	myLowUser = make(map[int64][]*UserRecommend, 0)
	userRecommendsMap = make(map[int64]*UserRecommend, 0)

	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		fmt.Println("今日分红错误用户获取失败2")
		return nil, err
	}

	for _, vUr := range userRecommends {
		userRecommendsMap[vUr.UserId] = vUr

		// 我的直推
		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(vUr.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if 0 >= myUserRecommendUserId {
			continue
		}

		if _, ok := myLowUser[myUserRecommendUserId]; !ok {
			myLowUser[myUserRecommendUserId] = make([]*UserRecommend, 0)
		}

		myLowUser[myUserRecommendUserId] = append(myLowUser[myUserRecommendUserId], vUr)
	}

	totalOne := float64(0)
	for _, v := range userReward1 {
		tmp := v.AmountUsdt * level1

		stop := false
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*1.5 {
			tmp = math.Abs(v.AmountUsdt*1.5 - v.AmountUsdtGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000

		if 0 >= tmp {
			continue
		}
		totalOne += tmp
		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserReward(ctx, v.ID, tmp, v.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, v)
			}

			if stop {
				stopUserIds[v.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
						if 0 >= len(vTmpRecommendUserIds) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, v.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	for _, v := range userReward2 {
		tmp := v.AmountUsdt * level2

		stop := false
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*1.8 {
			tmp = math.Abs(v.AmountUsdt*1.8 - v.AmountUsdtGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000
		if 0 >= tmp {
			continue
		}

		totalOne += tmp
		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserReward(ctx, v.ID, tmp, v.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, v)
			}

			if stop {
				stopUserIds[v.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
						if 0 >= len(vTmpRecommendUserIds) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, v.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	for _, v := range userReward3 {
		tmp := v.AmountUsdt * level3

		stop := false
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2 {
			tmp = math.Abs(v.AmountUsdt*2 - v.AmountUsdtGet)
			stop = true
		}

		tmp = math.Round(tmp*10000000) / 10000000
		if 0 >= tmp {
			continue
		}

		totalOne += tmp
		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserReward(ctx, v.ID, tmp, v.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, v)
			}

			if stop {
				stopUserIds[v.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
						if 0 >= len(vTmpRecommendUserIds) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, v.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	for _, v := range userReward4 {
		tmp := v.AmountUsdt * level4

		stop := false
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2.3 {
			tmp = math.Abs(v.AmountUsdt*2.3 - v.AmountUsdtGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000
		if 0 >= tmp {
			continue
		}

		totalOne += tmp
		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserReward(ctx, v.ID, tmp, v.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, v)
			}

			if stop {
				stopUserIds[v.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
						if 0 >= len(vTmpRecommendUserIds) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, v.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	for _, v := range userReward5 {
		tmp := v.AmountUsdt * level5

		stop := false
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2.6 {
			tmp = math.Abs(v.AmountUsdt*2.6 - v.AmountUsdtGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000
		if 0 >= tmp {
			continue
		}
		totalOne += tmp
		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserReward(ctx, v.ID, tmp, v.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, v)
			}

			if stop {
				stopUserIds[v.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
						if 0 >= len(vTmpRecommendUserIds) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, v.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	for _, v := range userReward6 {
		tmp := v.AmountUsdt * level6

		stop := false
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*3 {
			tmp = math.Abs(v.AmountUsdt*3 - v.AmountUsdtGet)
			stop = true
		}
		tmp = math.Round(tmp*10000000) / 10000000
		if 0 >= tmp {
			continue
		}

		totalOne += tmp
		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)
			code, err = uuc.uiRepo.UpdateUserReward(ctx, v.ID, tmp, v.AmountUsdt, stop)
			if code > 0 && err != nil {
				fmt.Println("错误分红静态：", err, v)
			}

			if stop {
				stopUserIds[v.ID] = true // 出局

				if nil != userRecommend && "" != userRecommend.RecommendCode {
					var tmpRecommendUserIds []string
					tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
					for _, vTmpRecommendUserIds := range tmpRecommendUserIds {
						if 0 >= len(vTmpRecommendUserIds) {
							continue
						}

						myUserRecommendUserId, _ := strconv.ParseInt(vTmpRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendUserId, v.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红静态：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily", err, v)
			continue
		}
	}

	// 更新总数
	//err = uuc.uiRepo.UpdateTotalOne(ctx, totalOne)
	//if err != nil {
	//	fmt.Println("更新总数：", err, totalOne)
	//}

	// 直推
	for _, v := range userReward1 {
		if 1 == v.LockReward {
			continue
		}
		tmp := v.AmountUsdt * level1
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*1.5 {
			tmp = math.Abs(v.AmountUsdt*1.5 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		// 超过金额
		if v.AmountRecommendUsdtGet >= v.AmountUsdt {
			continue
		}

		// 直推奖
		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if _, ok := usersMap[myUserRecommendUserId]; !ok {
			fmt.Println("错误分红直推，信息缺失,user：", err, v)
			continue
		}

		tmpRecommendUser := usersMap[myUserRecommendUserId]
		if nil == tmpRecommendUser {
			fmt.Println("错误分红直推，信息缺失,user1：", err, v)
			continue
		}

		// 出局的
		if 0 >= tmpRecommendUser.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[myUserRecommendUserId]; ok {
			continue
		}

		var (
			stopRecommend bool
			num           float64
		)
		if 1 == tmpRecommendUser.Last {
			num = 1.5
		} else if 2 == tmpRecommendUser.Last {
			num = 1.8
		} else if 3 == tmpRecommendUser.Last {
			num = 2
		} else if 4 == tmpRecommendUser.Last {
			num = 2.3
		} else if 5 == tmpRecommendUser.Last {
			num = 2.6
		} else if 6 == tmpRecommendUser.Last {
			num = 3
		} else {
			continue
		}

		tmpRecommendAmount := tmp
		enough := false
		if tmpRecommendAmount+v.AmountRecommendUsdtGet >= v.AmountUsdt {
			tmpRecommendAmount = math.Abs(v.AmountUsdt - v.AmountRecommendUsdtGet)
			enough = true
		}

		if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
			tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
			stopRecommend = true
		}

		tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
		if 0 >= tmpRecommendAmount {
			continue
		}
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardRecommend(ctx, myUserRecommendUserId, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, v.Address)
			if code > 0 && err != nil {
				fmt.Println("错误分红直推：", err, tmpRecommendUser)
			}

			// 修改已获得直推金额
			if _, ok := stopUserIds[v.ID]; !ok {
				err = uuc.uiRepo.UpdateUserRewardRecommendUserGet(ctx, v.ID, tmpRecommendAmount, enough, v.AmountUsdt)
				if err != nil {
					fmt.Println("错误分红直推：", err, tmpRecommendUser)
				}
			}

			if stopRecommend {
				stopUserIds[myUserRecommendUserId] = true // 出局

				// 推荐人
				var (
					userRecommendRecommend *UserRecommend
				)
				if _, ok := userRecommendsMap[myUserRecommendUserId]; ok {
					userRecommendRecommend = userRecommendsMap[myUserRecommendUserId]
				} else {
					fmt.Println("错误分红直推，信息缺失：", err, v)
				}

				if nil != userRecommendRecommend && "" != userRecommendRecommend.RecommendCode {
					var tmpRecommendRecommendUserIds []string
					tmpRecommendRecommendUserIds = strings.Split(userRecommendRecommend.RecommendCode, "D")

					for _, vTmpRecommendRecommendUserIds := range tmpRecommendRecommendUserIds {
						if 0 >= len(vTmpRecommendRecommendUserIds) {
							continue
						}

						myUserRecommendRecommendUserId, _ := strconv.ParseInt(vTmpRecommendRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendRecommendUserId, tmpRecommendUser.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红直推：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily recommend", err, v)
			continue
		}
	}

	for _, v := range userReward2 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level2
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*1.8 {
			tmp = math.Abs(v.AmountUsdt*1.8 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		// 超过金额
		if v.AmountRecommendUsdtGet >= v.AmountUsdt {
			continue
		}

		// 直推奖
		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if _, ok := usersMap[myUserRecommendUserId]; !ok {
			fmt.Println("错误分红直推，信息缺失,user：", err, v)
			continue
		}

		tmpRecommendUser := usersMap[myUserRecommendUserId]
		if nil == tmpRecommendUser {
			fmt.Println("错误分红直推，信息缺失,user1：", err, v)
			continue
		}

		// 出局的
		if 0 >= tmpRecommendUser.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[myUserRecommendUserId]; ok {
			continue
		}

		var (
			stopRecommend bool
			num           float64
		)
		if 1 == tmpRecommendUser.Last {
			num = 1.5
		} else if 2 == tmpRecommendUser.Last {
			num = 1.8
		} else if 3 == tmpRecommendUser.Last {
			num = 2
		} else if 4 == tmpRecommendUser.Last {
			num = 2.3
		} else if 5 == tmpRecommendUser.Last {
			num = 2.6
		} else if 6 == tmpRecommendUser.Last {
			num = 3
		} else {
			continue
		}

		tmpRecommendAmount := tmp
		enough := false
		if tmpRecommendAmount+v.AmountRecommendUsdtGet >= v.AmountUsdt {
			tmpRecommendAmount = math.Abs(v.AmountUsdt - v.AmountRecommendUsdtGet)
			enough = true
		}

		if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
			tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
			stopRecommend = true
		}

		tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
		if 0 >= tmpRecommendAmount {
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardRecommend(ctx, myUserRecommendUserId, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, v.Address)
			if code > 0 && err != nil {
				fmt.Println("错误分红直推：", err, tmpRecommendUser)
			}

			// 修改已获得直推金额
			if _, ok := stopUserIds[v.ID]; !ok {
				err = uuc.uiRepo.UpdateUserRewardRecommendUserGet(ctx, v.ID, tmpRecommendAmount, enough, v.AmountUsdt)
				if err != nil {
					fmt.Println("错误分红直推：", err, tmpRecommendUser)
				}
			}

			if stopRecommend {
				stopUserIds[myUserRecommendUserId] = true // 出局

				// 推荐人
				var (
					userRecommendRecommend *UserRecommend
				)
				if _, ok := userRecommendsMap[myUserRecommendUserId]; ok {
					userRecommendRecommend = userRecommendsMap[myUserRecommendUserId]
				} else {
					fmt.Println("错误分红直推，信息缺失：", err, v)
				}

				if nil != userRecommendRecommend && "" != userRecommendRecommend.RecommendCode {
					var tmpRecommendRecommendUserIds []string
					tmpRecommendRecommendUserIds = strings.Split(userRecommendRecommend.RecommendCode, "D")

					for _, vTmpRecommendRecommendUserIds := range tmpRecommendRecommendUserIds {
						if 0 >= len(vTmpRecommendRecommendUserIds) {
							continue
						}

						myUserRecommendRecommendUserId, _ := strconv.ParseInt(vTmpRecommendRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendRecommendUserId, tmpRecommendUser.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红直推：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily recommend", err, v)
			continue
		}
	}

	for _, v := range userReward3 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level3
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2 {
			tmp = math.Abs(v.AmountUsdt*2 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		// 超过金额
		if v.AmountRecommendUsdtGet >= v.AmountUsdt {
			continue
		}

		// 直推奖
		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if _, ok := usersMap[myUserRecommendUserId]; !ok {
			fmt.Println("错误分红直推，信息缺失,user：", err, v)
			continue
		}

		tmpRecommendUser := usersMap[myUserRecommendUserId]
		if nil == tmpRecommendUser {
			fmt.Println("错误分红直推，信息缺失,user1：", err, v)
			continue
		}

		// 出局的
		if 0 >= tmpRecommendUser.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[myUserRecommendUserId]; ok {
			continue
		}

		var (
			stopRecommend bool
			num           float64
		)
		if 1 == tmpRecommendUser.Last {
			num = 1.5
		} else if 2 == tmpRecommendUser.Last {
			num = 1.8
		} else if 3 == tmpRecommendUser.Last {
			num = 2
		} else if 4 == tmpRecommendUser.Last {
			num = 2.3
		} else if 5 == tmpRecommendUser.Last {
			num = 2.6
		} else if 6 == tmpRecommendUser.Last {
			num = 3
		} else {
			continue
		}

		tmpRecommendAmount := tmp
		enough := false
		if tmpRecommendAmount+v.AmountRecommendUsdtGet >= v.AmountUsdt {
			tmpRecommendAmount = math.Abs(v.AmountUsdt - v.AmountRecommendUsdtGet)
			enough = true
		}

		if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
			tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
			stopRecommend = true
		}

		tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
		if 0 >= tmpRecommendAmount {
			continue
		}
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardRecommend(ctx, myUserRecommendUserId, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, v.Address)
			if code > 0 && err != nil {
				fmt.Println("错误分红直推：", err, tmpRecommendUser)
			}

			// 修改已获得直推金额
			if _, ok := stopUserIds[v.ID]; !ok {
				err = uuc.uiRepo.UpdateUserRewardRecommendUserGet(ctx, v.ID, tmpRecommendAmount, enough, v.AmountUsdt)
				if err != nil {
					fmt.Println("错误分红直推：", err, tmpRecommendUser)
				}
			}

			if stopRecommend {
				stopUserIds[myUserRecommendUserId] = true // 出局

				// 推荐人
				var (
					userRecommendRecommend *UserRecommend
				)
				if _, ok := userRecommendsMap[myUserRecommendUserId]; ok {
					userRecommendRecommend = userRecommendsMap[myUserRecommendUserId]
				} else {
					fmt.Println("错误分红直推，信息缺失：", err, v)
				}

				if nil != userRecommendRecommend && "" != userRecommendRecommend.RecommendCode {
					var tmpRecommendRecommendUserIds []string
					tmpRecommendRecommendUserIds = strings.Split(userRecommendRecommend.RecommendCode, "D")

					for _, vTmpRecommendRecommendUserIds := range tmpRecommendRecommendUserIds {
						if 0 >= len(vTmpRecommendRecommendUserIds) {
							continue
						}

						myUserRecommendRecommendUserId, _ := strconv.ParseInt(vTmpRecommendRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendRecommendUserId, tmpRecommendUser.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红直推：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily recommend", err, v)
			continue
		}
	}

	for _, v := range userReward4 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level4
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2.3 {
			tmp = math.Abs(v.AmountUsdt*2.3 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		// 超过金额
		if v.AmountRecommendUsdtGet >= v.AmountUsdt {
			continue
		}

		// 直推奖
		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if _, ok := usersMap[myUserRecommendUserId]; !ok {
			fmt.Println("错误分红直推，信息缺失,user：", err, v)
			continue
		}

		tmpRecommendUser := usersMap[myUserRecommendUserId]
		if nil == tmpRecommendUser {
			fmt.Println("错误分红直推，信息缺失,user1：", err, v)
			continue
		}

		// 出局的
		if 0 >= tmpRecommendUser.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[myUserRecommendUserId]; ok {
			continue
		}

		var (
			stopRecommend bool
			num           float64
		)
		if 1 == tmpRecommendUser.Last {
			num = 1.5
		} else if 2 == tmpRecommendUser.Last {
			num = 1.8
		} else if 3 == tmpRecommendUser.Last {
			num = 2
		} else if 4 == tmpRecommendUser.Last {
			num = 2.3
		} else if 5 == tmpRecommendUser.Last {
			num = 2.6
		} else if 6 == tmpRecommendUser.Last {
			num = 3
		} else {
			continue
		}

		tmpRecommendAmount := tmp
		enough := false
		if tmpRecommendAmount+v.AmountRecommendUsdtGet >= v.AmountUsdt {
			tmpRecommendAmount = math.Abs(v.AmountUsdt - v.AmountRecommendUsdtGet)
			enough = true
		}

		if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
			tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
			stopRecommend = true
		}

		tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
		if 0 >= tmpRecommendAmount {
			continue
		}
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardRecommend(ctx, myUserRecommendUserId, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, v.Address)
			if code > 0 && err != nil {
				fmt.Println("错误分红直推：", err, tmpRecommendUser)
			}

			// 修改已获得直推金额
			if _, ok := stopUserIds[v.ID]; !ok {
				err = uuc.uiRepo.UpdateUserRewardRecommendUserGet(ctx, v.ID, tmpRecommendAmount, enough, v.AmountUsdt)
				if err != nil {
					fmt.Println("错误分红直推：", err, tmpRecommendUser)
				}
			}

			if stopRecommend {
				stopUserIds[myUserRecommendUserId] = true // 出局

				// 推荐人
				var (
					userRecommendRecommend *UserRecommend
				)
				if _, ok := userRecommendsMap[myUserRecommendUserId]; ok {
					userRecommendRecommend = userRecommendsMap[myUserRecommendUserId]
				} else {
					fmt.Println("错误分红直推，信息缺失：", err, v)
				}

				if nil != userRecommendRecommend && "" != userRecommendRecommend.RecommendCode {
					var tmpRecommendRecommendUserIds []string
					tmpRecommendRecommendUserIds = strings.Split(userRecommendRecommend.RecommendCode, "D")

					for _, vTmpRecommendRecommendUserIds := range tmpRecommendRecommendUserIds {
						if 0 >= len(vTmpRecommendRecommendUserIds) {
							continue
						}

						myUserRecommendRecommendUserId, _ := strconv.ParseInt(vTmpRecommendRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendRecommendUserId, tmpRecommendUser.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红直推：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily recommend", err, v)
			continue
		}
	}

	for _, v := range userReward5 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level5
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2.6 {
			tmp = math.Abs(v.AmountUsdt*2.6 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		// 超过金额
		if v.AmountRecommendUsdtGet >= v.AmountUsdt {
			continue
		}

		// 直推奖
		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if _, ok := usersMap[myUserRecommendUserId]; !ok {
			fmt.Println("错误分红直推，信息缺失,user：", err, v)
			continue
		}

		tmpRecommendUser := usersMap[myUserRecommendUserId]
		if nil == tmpRecommendUser {
			fmt.Println("错误分红直推，信息缺失,user1：", err, v)
			continue
		}

		// 出局的
		if 0 >= tmpRecommendUser.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[myUserRecommendUserId]; ok {
			continue
		}

		var (
			stopRecommend bool
			num           float64
		)
		if 1 == tmpRecommendUser.Last {
			num = 1.5
		} else if 2 == tmpRecommendUser.Last {
			num = 1.8
		} else if 3 == tmpRecommendUser.Last {
			num = 2
		} else if 4 == tmpRecommendUser.Last {
			num = 2.3
		} else if 5 == tmpRecommendUser.Last {
			num = 2.6
		} else if 6 == tmpRecommendUser.Last {
			num = 3
		} else {
			continue
		}

		tmpRecommendAmount := tmp
		enough := false
		if tmpRecommendAmount+v.AmountRecommendUsdtGet >= v.AmountUsdt {
			tmpRecommendAmount = math.Abs(v.AmountUsdt - v.AmountRecommendUsdtGet)
			enough = true
		}

		if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
			tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
			stopRecommend = true
		}

		tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
		if 0 >= tmpRecommendAmount {
			continue
		}
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardRecommend(ctx, myUserRecommendUserId, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, v.Address)
			if code > 0 && err != nil {
				fmt.Println("错误分红直推：", err, tmpRecommendUser)
			}

			// 修改已获得直推金额
			if _, ok := stopUserIds[v.ID]; !ok {
				err = uuc.uiRepo.UpdateUserRewardRecommendUserGet(ctx, v.ID, tmpRecommendAmount, enough, v.AmountUsdt)
				if err != nil {
					fmt.Println("错误分红直推：", err, tmpRecommendUser)
				}
			}

			if stopRecommend {
				stopUserIds[myUserRecommendUserId] = true // 出局

				// 推荐人
				var (
					userRecommendRecommend *UserRecommend
				)
				if _, ok := userRecommendsMap[myUserRecommendUserId]; ok {
					userRecommendRecommend = userRecommendsMap[myUserRecommendUserId]
				} else {
					fmt.Println("错误分红直推，信息缺失：", err, v)
				}

				if nil != userRecommendRecommend && "" != userRecommendRecommend.RecommendCode {
					var tmpRecommendRecommendUserIds []string
					tmpRecommendRecommendUserIds = strings.Split(userRecommendRecommend.RecommendCode, "D")

					for _, vTmpRecommendRecommendUserIds := range tmpRecommendRecommendUserIds {
						if 0 >= len(vTmpRecommendRecommendUserIds) {
							continue
						}

						myUserRecommendRecommendUserId, _ := strconv.ParseInt(vTmpRecommendRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendRecommendUserId, tmpRecommendUser.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红直推：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily recommend", err, v)
			continue
		}
	}

	for _, v := range userReward6 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level6
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*3 {
			tmp = math.Abs(v.AmountUsdt*3 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红静态，信息缺失：", err, v)
		}

		// 超过金额
		if v.AmountRecommendUsdtGet >= v.AmountUsdt {
			continue
		}

		// 直推奖
		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			myUserRecommendUserId int64
			tmpRecommendUserIds   []string
		)

		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
		if 2 <= len(tmpRecommendUserIds) {
			myUserRecommendUserId, _ = strconv.ParseInt(tmpRecommendUserIds[len(tmpRecommendUserIds)-1], 10, 64) // 最后一位是直推人
		}

		if _, ok := usersMap[myUserRecommendUserId]; !ok {
			fmt.Println("错误分红直推，信息缺失,user：", err, v)
			continue
		}

		tmpRecommendUser := usersMap[myUserRecommendUserId]
		if nil == tmpRecommendUser {
			fmt.Println("错误分红直推，信息缺失,user1：", err, v)
			continue
		}

		// 出局的
		if 0 >= tmpRecommendUser.AmountUsdt {
			continue
		}

		// 本次执行已经出局
		if _, ok := stopUserIds[myUserRecommendUserId]; ok {
			continue
		}

		var (
			stopRecommend bool
			num           float64
		)
		if 1 == tmpRecommendUser.Last {
			num = 1.5
		} else if 2 == tmpRecommendUser.Last {
			num = 1.8
		} else if 3 == tmpRecommendUser.Last {
			num = 2
		} else if 4 == tmpRecommendUser.Last {
			num = 2.3
		} else if 5 == tmpRecommendUser.Last {
			num = 2.6
		} else if 6 == tmpRecommendUser.Last {
			num = 3
		} else {
			continue
		}

		tmpRecommendAmount := tmp
		enough := false
		if tmpRecommendAmount+v.AmountRecommendUsdtGet >= v.AmountUsdt {
			tmpRecommendAmount = math.Abs(v.AmountUsdt - v.AmountRecommendUsdtGet)
			enough = true
		}

		if tmpRecommendAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
			tmpRecommendAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
			stopRecommend = true
		}

		tmpRecommendAmount = math.Round(tmpRecommendAmount*10000000) / 10000000
		if 0 >= tmpRecommendAmount {
			continue
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardRecommend(ctx, myUserRecommendUserId, tmpRecommendAmount, tmpRecommendUser.AmountUsdt, stopRecommend, v.Address)
			if code > 0 && err != nil {
				fmt.Println("错误分红直推：", err, tmpRecommendUser)
			}

			// 修改已获得直推金额
			if _, ok := stopUserIds[v.ID]; !ok {
				err = uuc.uiRepo.UpdateUserRewardRecommendUserGet(ctx, v.ID, tmpRecommendAmount, enough, v.AmountUsdt)
				if err != nil {
					fmt.Println("错误分红直推：", err, tmpRecommendUser)
				}
			}

			if stopRecommend {
				stopUserIds[myUserRecommendUserId] = true // 出局

				// 推荐人
				var (
					userRecommendRecommend *UserRecommend
				)
				if _, ok := userRecommendsMap[myUserRecommendUserId]; ok {
					userRecommendRecommend = userRecommendsMap[myUserRecommendUserId]
				} else {
					fmt.Println("错误分红直推，信息缺失：", err, v)
				}

				if nil != userRecommendRecommend && "" != userRecommendRecommend.RecommendCode {
					var tmpRecommendRecommendUserIds []string
					tmpRecommendRecommendUserIds = strings.Split(userRecommendRecommend.RecommendCode, "D")

					for _, vTmpRecommendRecommendUserIds := range tmpRecommendRecommendUserIds {
						if 0 >= len(vTmpRecommendRecommendUserIds) {
							continue
						}

						myUserRecommendRecommendUserId, _ := strconv.ParseInt(vTmpRecommendRecommendUserIds, 10, 64) // 最后一位是直推人
						if 0 >= myUserRecommendRecommendUserId {
							continue
						}

						// 减掉业绩
						err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendRecommendUserId, tmpRecommendUser.AmountUsdt)
						if err != nil {
							fmt.Println("错误分红直推：", err, v)
						}
					}
				}
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily recommend", err, v)
			continue
		}
	}

	// 大小区
	for _, v := range userReward1 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level1
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*1.5 {
			tmp = math.Abs(v.AmountUsdt*1.5 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红小区，信息缺失：", err, v)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, v)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, v)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].AmountUsdt {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 6 == tmpRecommendUser.Vip {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 7 == tmpRecommendUser.Vip {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 8 == tmpRecommendUser.Vip {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000 <= tmpAreaMin && 5000 > tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 5000 <= tmpAreaMin && 30000 > tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 30000 <= tmpAreaMin && 100000 > tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 100000 <= tmpAreaMin && 300000 > tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 300000 <= tmpAreaMin && 1000000 > tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 1000000 <= tmpAreaMin && 3000000 > tmpAreaMin {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 3000000 <= tmpAreaMin && 10000000 > tmpAreaMin {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 10000000 <= tmpAreaMin {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			}

			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmp = tmp * v0
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			tmpAreaAmount := tmp
			var (
				stopArea bool
				num      float64
			)
			if 1 == tmpRecommendUser.Last {
				num = 1.5
			} else if 2 == tmpRecommendUser.Last {
				num = 1.8
			} else if 3 == tmpRecommendUser.Last {
				num = 2
			} else if 4 == tmpRecommendUser.Last {
				num = 2.3
			} else if 5 == tmpRecommendUser.Last {
				num = 2.6
			} else if 6 == tmpRecommendUser.Last {
				num = 3
			} else {
				continue
			}

			if tmpAreaAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
				tmpAreaAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
				stopArea = true
			}

			// 平级，结束
			tmpLevel := false
			if 0 < currentLevel {
				if currentLevel == lastLevel {
					tmpLevel = true
				}
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpAreaAmount, tmpRecommendUser.AmountUsdt, tmpLevel, stopArea, int64(currentLevel), int64(i), v.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}

	}

	for _, v := range userReward2 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level2
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*1.8 {
			tmp = math.Abs(v.AmountUsdt*1.8 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红小区，信息缺失：", err, v)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, v)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, v)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].AmountUsdt {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 6 == tmpRecommendUser.Vip {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 7 == tmpRecommendUser.Vip {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 8 == tmpRecommendUser.Vip {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000 <= tmpAreaMin && 5000 > tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 5000 <= tmpAreaMin && 30000 > tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 30000 <= tmpAreaMin && 100000 > tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 100000 <= tmpAreaMin && 300000 > tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 300000 <= tmpAreaMin && 1000000 > tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 1000000 <= tmpAreaMin && 3000000 > tmpAreaMin {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 3000000 <= tmpAreaMin && 10000000 > tmpAreaMin {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 10000000 <= tmpAreaMin {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			}

			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmp = tmp * v0
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			tmpAreaAmount := tmp
			var (
				stopArea bool
				num      float64
			)
			if 1 == tmpRecommendUser.Last {
				num = 1.5
			} else if 2 == tmpRecommendUser.Last {
				num = 1.8
			} else if 3 == tmpRecommendUser.Last {
				num = 2
			} else if 4 == tmpRecommendUser.Last {
				num = 2.3
			} else if 5 == tmpRecommendUser.Last {
				num = 2.6
			} else if 6 == tmpRecommendUser.Last {
				num = 3
			} else {
				continue
			}

			if tmpAreaAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
				tmpAreaAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
				stopArea = true
			}

			// 平级，结束
			tmpLevel := false
			if 0 < currentLevel {
				if currentLevel == lastLevel {
					tmpLevel = true
				}
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpAreaAmount, tmpRecommendUser.AmountUsdt, tmpLevel, stopArea, int64(currentLevel), int64(i), v.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}

	}

	for _, v := range userReward3 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level3
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2 {
			tmp = math.Abs(v.AmountUsdt*2 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红小区，信息缺失：", err, v)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, v)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, v)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].AmountUsdt {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 6 == tmpRecommendUser.Vip {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 7 == tmpRecommendUser.Vip {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 8 == tmpRecommendUser.Vip {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000 <= tmpAreaMin && 5000 > tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 5000 <= tmpAreaMin && 30000 > tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 30000 <= tmpAreaMin && 100000 > tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 100000 <= tmpAreaMin && 300000 > tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 300000 <= tmpAreaMin && 1000000 > tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 1000000 <= tmpAreaMin && 3000000 > tmpAreaMin {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 3000000 <= tmpAreaMin && 10000000 > tmpAreaMin {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 10000000 <= tmpAreaMin {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			}

			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmp = tmp * v0
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			tmpAreaAmount := tmp
			var (
				stopArea bool
				num      float64
			)
			if 1 == tmpRecommendUser.Last {
				num = 1.5
			} else if 2 == tmpRecommendUser.Last {
				num = 1.8
			} else if 3 == tmpRecommendUser.Last {
				num = 2
			} else if 4 == tmpRecommendUser.Last {
				num = 2.3
			} else if 5 == tmpRecommendUser.Last {
				num = 2.6
			} else if 6 == tmpRecommendUser.Last {
				num = 3
			} else {
				continue
			}

			if tmpAreaAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
				tmpAreaAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
				stopArea = true
			}

			// 平级，结束
			tmpLevel := false
			if 0 < currentLevel {
				if currentLevel == lastLevel {
					tmpLevel = true
				}
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpAreaAmount, tmpRecommendUser.AmountUsdt, tmpLevel, stopArea, int64(currentLevel), int64(i), v.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}

	}

	for _, v := range userReward4 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level4
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2.3 {
			tmp = math.Abs(v.AmountUsdt*2.3 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红小区，信息缺失：", err, v)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, v)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, v)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].AmountUsdt {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 6 == tmpRecommendUser.Vip {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 7 == tmpRecommendUser.Vip {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 8 == tmpRecommendUser.Vip {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000 <= tmpAreaMin && 5000 > tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 5000 <= tmpAreaMin && 30000 > tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 30000 <= tmpAreaMin && 100000 > tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 100000 <= tmpAreaMin && 300000 > tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 300000 <= tmpAreaMin && 1000000 > tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 1000000 <= tmpAreaMin && 3000000 > tmpAreaMin {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 3000000 <= tmpAreaMin && 10000000 > tmpAreaMin {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 10000000 <= tmpAreaMin {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			}

			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmp = tmp * v0
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			tmpAreaAmount := tmp
			var (
				stopArea bool
				num      float64
			)
			if 1 == tmpRecommendUser.Last {
				num = 1.5
			} else if 2 == tmpRecommendUser.Last {
				num = 1.8
			} else if 3 == tmpRecommendUser.Last {
				num = 2
			} else if 4 == tmpRecommendUser.Last {
				num = 2.3
			} else if 5 == tmpRecommendUser.Last {
				num = 2.6
			} else if 6 == tmpRecommendUser.Last {
				num = 3
			} else {
				continue
			}

			if tmpAreaAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
				tmpAreaAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
				stopArea = true
			}

			// 平级，结束
			tmpLevel := false
			if 0 < currentLevel {
				if currentLevel == lastLevel {
					tmpLevel = true
				}
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpAreaAmount, tmpRecommendUser.AmountUsdt, tmpLevel, stopArea, int64(currentLevel), int64(i), v.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}

	}

	for _, v := range userReward5 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level5
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*2.6 {
			tmp = math.Abs(v.AmountUsdt*2.6 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红小区，信息缺失：", err, v)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, v)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, v)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].AmountUsdt {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 6 == tmpRecommendUser.Vip {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 7 == tmpRecommendUser.Vip {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 8 == tmpRecommendUser.Vip {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000 <= tmpAreaMin && 5000 > tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 5000 <= tmpAreaMin && 30000 > tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 30000 <= tmpAreaMin && 100000 > tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 100000 <= tmpAreaMin && 300000 > tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 300000 <= tmpAreaMin && 1000000 > tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 1000000 <= tmpAreaMin && 3000000 > tmpAreaMin {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 3000000 <= tmpAreaMin && 10000000 > tmpAreaMin {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 10000000 <= tmpAreaMin {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			}

			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmp = tmp * v0
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			tmpAreaAmount := tmp
			var (
				stopArea bool
				num      float64
			)
			if 1 == tmpRecommendUser.Last {
				num = 1.5
			} else if 2 == tmpRecommendUser.Last {
				num = 1.8
			} else if 3 == tmpRecommendUser.Last {
				num = 2
			} else if 4 == tmpRecommendUser.Last {
				num = 2.3
			} else if 5 == tmpRecommendUser.Last {
				num = 2.6
			} else if 6 == tmpRecommendUser.Last {
				num = 3
			} else {
				continue
			}

			if tmpAreaAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
				tmpAreaAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
				stopArea = true
			}

			// 平级，结束
			tmpLevel := false
			if 0 < currentLevel {
				if currentLevel == lastLevel {
					tmpLevel = true
				}
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpAreaAmount, tmpRecommendUser.AmountUsdt, tmpLevel, stopArea, int64(currentLevel), int64(i), v.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}

	}

	for _, v := range userReward6 {
		if 1 == v.LockReward {
			continue
		}

		tmp := v.AmountUsdt * level6
		if tmp+v.AmountUsdtGet >= v.AmountUsdt*3 {
			tmp = math.Abs(v.AmountUsdt*3 - v.AmountUsdtGet)
		}

		if 0 >= tmp {
			continue
		}

		// 推荐人
		var (
			userRecommend *UserRecommend
		)
		if _, ok := userRecommendsMap[v.ID]; ok {
			userRecommend = userRecommendsMap[v.ID]
		} else {
			fmt.Println("错误分红小区，信息缺失：", err, v)
		}

		if nil == userRecommend || "" == userRecommend.RecommendCode {
			continue
		}

		var (
			tmpRecommendUserIds []string
		)
		tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")

		lastLevel := 0
		lastLevelNum := float64(0)
		for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
			currentLevel := 0
			tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
			if 0 >= tmpUserId {
				continue
			}

			// 本次执行已经出局
			if _, ok := stopUserIds[tmpUserId]; ok {
				continue
			}

			if _, ok := usersMap[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失,user：", err, v)
				continue
			}

			tmpRecommendUser := usersMap[tmpUserId]
			if nil == tmpRecommendUser {
				fmt.Println("错误分红小区，信息缺失,user1：", err, v)
				continue
			}

			if 0 >= tmpRecommendUser.AmountUsdt {
				continue
			}

			// 我的下级
			if _, ok := myLowUser[tmpUserId]; !ok {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 0 >= len(myLowUser[tmpUserId]) {
				fmt.Println("错误分红小区，信息缺失3：", err, tmpUserId, v)
				continue
			}

			if 1 >= len(myLowUser[tmpUserId]) {
				continue
			}

			// 获取业绩
			tmpAreaMax := float64(0)
			tmpMaxId := int64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if _, ok := usersMap[vMyLowUser.UserId]; !ok {
					fmt.Println("错误分红小区，信息缺失4：", err, tmpUserId, v)
					continue
				}

				if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount+usersMap[vMyLowUser.UserId].AmountUsdt {
					tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
					tmpMaxId = vMyLowUser.UserId
				}
			}

			if 0 >= tmpMaxId {
				continue
			}

			tmpAreaMin := float64(0)
			for _, vMyLowUser := range myLowUser[tmpUserId] {
				if tmpMaxId != vMyLowUser.UserId {
					tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount + usersMap[vMyLowUser.UserId].AmountUsdt
				}
			}

			tmpLastLevelNum := float64(0)
			if 0 < tmpRecommendUser.Vip {
				if 1 == tmpRecommendUser.Vip {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 2 == tmpRecommendUser.Vip {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 3 == tmpRecommendUser.Vip {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 4 == tmpRecommendUser.Vip {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 5 == tmpRecommendUser.Vip {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 6 == tmpRecommendUser.Vip {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 7 == tmpRecommendUser.Vip {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 8 == tmpRecommendUser.Vip {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			} else {
				if 1000 <= tmpAreaMin && 5000 > tmpAreaMin {
					currentLevel = 1
					tmpLastLevelNum = vv1
				} else if 5000 <= tmpAreaMin && 30000 > tmpAreaMin {
					currentLevel = 2
					tmpLastLevelNum = v2
				} else if 30000 <= tmpAreaMin && 100000 > tmpAreaMin {
					currentLevel = 3
					tmpLastLevelNum = v3
				} else if 100000 <= tmpAreaMin && 300000 > tmpAreaMin {
					currentLevel = 4
					tmpLastLevelNum = v4
				} else if 300000 <= tmpAreaMin && 1000000 > tmpAreaMin {
					currentLevel = 5
					tmpLastLevelNum = v5
				} else if 1000000 <= tmpAreaMin && 3000000 > tmpAreaMin {
					currentLevel = 6
					tmpLastLevelNum = v6
				} else if 3000000 <= tmpAreaMin && 10000000 > tmpAreaMin {
					currentLevel = 7
					tmpLastLevelNum = v7
				} else if 10000000 <= tmpAreaMin {
					currentLevel = 8
					tmpLastLevelNum = v8
				} else {
					// 跳过，没级别
					continue
				}
			}

			// 级别低跳过
			if currentLevel < lastLevel {
				continue
			} else if currentLevel == lastLevel {
				tmp = tmp * v0
			} else {
				// 级差
				if tmpLastLevelNum < lastLevelNum {
					fmt.Println("错误分红小区，配置，信息缺错误：", err, tmpUserId, v, tmpLastLevelNum, lastLevelNum)
					continue
				}

				tmp = tmp * (tmpLastLevelNum - lastLevelNum)
			}

			tmpAreaAmount := tmp
			var (
				stopArea bool
				num      float64
			)
			if 1 == tmpRecommendUser.Last {
				num = 1.5
			} else if 2 == tmpRecommendUser.Last {
				num = 1.8
			} else if 3 == tmpRecommendUser.Last {
				num = 2
			} else if 4 == tmpRecommendUser.Last {
				num = 2.3
			} else if 5 == tmpRecommendUser.Last {
				num = 2.6
			} else if 6 == tmpRecommendUser.Last {
				num = 3
			} else {
				continue
			}

			if tmpAreaAmount+tmpRecommendUser.AmountUsdtGet >= tmpRecommendUser.AmountUsdt*num {
				tmpAreaAmount = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
				stopArea = true
			}

			// 平级，结束
			tmpLevel := false
			if 0 < currentLevel {
				if currentLevel == lastLevel {
					tmpLevel = true
				}
			}

			// 分红
			tmpAreaAmount = math.Round(tmpAreaAmount*10000000) / 10000000
			if 0 >= tmpAreaAmount {
				continue
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var (
					code int64
				)

				code, err = uuc.uiRepo.UpdateUserRewardArea(ctx, tmpRecommendUser.ID, tmpAreaAmount, tmpRecommendUser.AmountUsdt, tmpLevel, stopArea, int64(currentLevel), int64(i), v.Address)
				if code > 0 && err != nil {
					fmt.Println("错误分红小区：", err, tmpRecommendUser)
				}

				if stopArea {
					stopUserIds[tmpRecommendUser.ID] = true // 出局

					// 推荐人
					var (
						userRecommendArea *UserRecommend
					)
					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
					} else {
						fmt.Println("错误分红小区，信息缺失7：", err, v)
					}

					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
						var tmpRecommendAreaUserIds []string
						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")

						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
							if 0 >= len(vTmpRecommendAreaUserIds) {
								continue
							}

							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
							if 0 >= myUserRecommendAreaUserId {
								continue
							}

							// 减掉业绩
							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
							if err != nil {
								fmt.Println("错误分红小区：", err, v)
							}
						}
					}
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily area", err, v)
			}

			// 平级，结束
			if tmpLevel {
				break
			}

			if currentLevel > lastLevel {
				lastLevel = currentLevel
				lastLevelNum = tmpLastLevelNum
			}
		}

	}

	// 社区奖励
	var (
		exchanges            []*UserBalanceRecord
		totalExchange        float64
		totalExchangeRate    float64
		totalExchangeRateTwo float64
	)
	exchanges, err = uuc.ubRepo.GetSystemYesterdayLocationReward(ctx, -1)
	if nil != err {
		return nil, nil
	}

	for _, v := range exchanges {
		totalExchange += v.AmountNewTwo
	}
	totalExchangeRate = totalExchange * 0.4
	totalExchangeRateTwo = totalExchange * 0.6

	fmt.Println("今日发放兑换：", totalExchange, totalExchangeRate, totalExchangeRateTwo)

	if 0 >= totalExchange {
		return nil, nil
	}

	var (
		stake         []*Stake
		stakeTotal    float64
		stakeTotalTwo float64
	)
	stake, err = uuc.ubRepo.GetStake(ctx)
	if nil != err {
		return nil, err
	}

	stakeOne := make([]*Stake, 0)
	stakeOneRemove := make([]*Stake, 0)
	stakeTwo := make([]*Stake, 0)
	stakeTwoRemove := make([]*Stake, 0)
	for _, v := range stake {
		if 0 != v.Status {
			continue
		}

		if 10 == v.Day {
			if v.CreatedAt.Add(10 * 24 * time.Hour).Before(time.Now()) {
				stakeOneRemove = append(stakeOneRemove, v)
				continue
			}

			stakeTotal += v.Amount
			stakeOne = append(stakeOne, v)
		}

		if 30 == v.Day {
			if v.CreatedAt.Add(30 * 24 * time.Hour).Before(time.Now()) {
				stakeTwoRemove = append(stakeTwoRemove, v)
				continue
			}

			stakeTotalTwo += v.Amount
			stakeTwo = append(stakeTwo, v)
		}
	}

	for _, v := range stakeOne {
		tmpStakeAmount := math.Round(v.Amount/stakeTotal*totalExchangeRate*10000000) / 10000000

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardStake(ctx, v.UserId, tmpStakeAmount, v.ID)
			if code > 0 && err != nil {
				fmt.Println("错误stake分红1：", err, v, tmpStakeAmount)
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily stake", err, v, tmpStakeAmount)
		}
	}

	for _, v := range stakeTwo {
		tmpStakeAmount := math.Round(v.Amount/stakeTotal*totalExchangeRateTwo*10000000) / 10000000

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardStake(ctx, v.UserId, tmpStakeAmount, v.ID)
			if code > 0 && err != nil {
				fmt.Println("错误stake分红1：", err, v, tmpStakeAmount)
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily stake", err, v, tmpStakeAmount)
		}
	}

	for _, v := range stakeOneRemove {
		tmpStakeAmount := math.Round((v.Amount+v.Reward)*10000000) / 10000000

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardStakeReomve(ctx, v.UserId, tmpStakeAmount, v.ID)
			if code > 0 && err != nil {
				fmt.Println("错误stake分红1：", err, v, tmpStakeAmount)
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily stake", err, v, tmpStakeAmount)
		}
	}

	for _, v := range stakeTwoRemove {
		tmpStakeAmount := math.Round((v.Amount+v.Reward)*10000000) / 10000000

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			var (
				code int64
			)

			code, err = uuc.uiRepo.UpdateUserRewardStakeReomve(ctx, v.UserId, tmpStakeAmount, v.ID)
			if code > 0 && err != nil {
				fmt.Println("错误stake分红1：", err, v, tmpStakeAmount)
				return err
			}

			return nil
		}); nil != err {
			fmt.Println("err reward daily stake", err, v, tmpStakeAmount)
		}
	}
	//
	//for _, vBuys := range buys {
	//	userId := vBuys.UserId
	//
	//	// 推荐人
	//	var (
	//		userRecommend *UserRecommend
	//	)
	//	if _, ok := userRecommendsMap[userId]; ok {
	//		userRecommend = userRecommendsMap[userId]
	//	} else {
	//		fmt.Println("错误分红社区，信息缺失：", err, vBuys)
	//	}
	//
	//	if nil == userRecommend || "" == userRecommend.RecommendCode {
	//		continue
	//	}
	//
	//	var (
	//		tmpRecommendUserIds []string
	//	)
	//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	//
	//	lastLevel := 0
	//	lastLevelNum := float64(0)
	//	for i := len(tmpRecommendUserIds) - 1; i >= 0; i-- {
	//		currentLevel := 0
	//
	//		tmpUserId, _ := strconv.ParseInt(tmpRecommendUserIds[i], 10, 64) // 最后一位是直推人
	//		if 0 >= tmpUserId {
	//			continue
	//		}
	//
	//		// 本次执行已经出局
	//		if _, ok := stopUserIds[tmpUserId]; ok {
	//			continue
	//		}
	//
	//		if _, ok := usersMap[tmpUserId]; !ok {
	//			fmt.Println("错误分红社区，信息缺失,user：", err, vBuys, tmpUserId)
	//			continue
	//		}
	//
	//		tmpRecommendUser := usersMap[tmpUserId]
	//		if nil == tmpRecommendUser {
	//			fmt.Println("错误分红社区，信息缺失,user1：", err, vBuys)
	//			continue
	//		}
	//
	//		// 我的下级
	//		if _, ok := myLowUser[tmpUserId]; !ok {
	//			fmt.Println("错误分红社区，信息缺失3：", err, tmpUserId, vBuys)
	//			continue
	//		}
	//
	//		if 0 >= len(myLowUser[tmpUserId]) {
	//			fmt.Println("错误分红社区，信息缺失3：", err, tmpUserId, vBuys)
	//			continue
	//		}
	//
	//		if 1 >= len(myLowUser[tmpUserId]) {
	//			continue
	//		}
	//
	//		// 获取业绩
	//		tmpAreaMax := float64(0)
	//		tmpMaxId := int64(0)
	//		for _, vMyLowUser := range myLowUser[tmpUserId] {
	//			if _, ok := usersMap[vMyLowUser.UserId]; !ok {
	//				fmt.Println("错误分红社区，信息缺失4：", err, tmpUserId, vBuys)
	//				continue
	//			}
	//
	//			if tmpAreaMax < usersMap[vMyLowUser.UserId].MyTotalAmount {
	//				tmpAreaMax = usersMap[vMyLowUser.UserId].MyTotalAmount
	//				tmpMaxId = vMyLowUser.UserId
	//			}
	//		}
	//
	//		if 0 >= tmpMaxId {
	//			continue
	//		}
	//
	//		tmpAreaMin := float64(0)
	//		for _, vMyLowUser := range myLowUser[tmpUserId] {
	//			if tmpMaxId != vMyLowUser.UserId {
	//				tmpAreaMin += usersMap[vMyLowUser.UserId].MyTotalAmount
	//			}
	//		}
	//
	//		tmpLastLevelNum := float64(0)
	//		if 100000 <= tmpAreaMin {
	//			currentLevel = 4
	//			tmpLastLevelNum = va4
	//		} else if 300000 <= tmpAreaMin {
	//			currentLevel = 5
	//			tmpLastLevelNum = va5
	//		} else if 1000000 <= tmpAreaMin {
	//			currentLevel = 6
	//			tmpLastLevelNum = va6
	//		} else if 3000000 <= tmpAreaMin {
	//			currentLevel = 7
	//			tmpLastLevelNum = va7
	//		} else if 10000000 <= tmpAreaMin {
	//			currentLevel = 8
	//			tmpLastLevelNum = va8
	//		} else {
	//			// 跳过，没级别
	//			continue
	//		}
	//
	//		// 级别低跳过
	//		if currentLevel <= lastLevel {
	//			if 8 == lastLevel {
	//				break
	//			}
	//
	//			continue
	//		} else {
	//			// 级差
	//			if tmpLastLevelNum < lastLevelNum {
	//				fmt.Println("错误分红社区，配置，信息缺错误：", err, tmpUserId, vBuys, tmpLastLevelNum, lastLevelNum)
	//				continue
	//			}
	//
	//			tmp := vBuys.AmountNew * (tmpLastLevelNum - lastLevelNum)
	//
	//			var (
	//				stopArea2 bool
	//				num       float64
	//			)
	//			if 1 == tmpRecommendUser.Last {
	//				num = 2
	//			} else if 2 == tmpRecommendUser.Last {
	//				num = 2.3
	//			} else if 3 == tmpRecommendUser.Last {
	//				num = 2.6
	//			} else if 4 == tmpRecommendUser.Last {
	//				num = 3
	//			} else {
	//				continue
	//			}
	//
	//			if !lessThanOrEqualZero(tmp+tmpRecommendUser.AmountUsdtGet, tmpRecommendUser.AmountUsdt*num, 1e-7) {
	//				tmp = math.Abs(tmpRecommendUser.AmountUsdt*num - tmpRecommendUser.AmountUsdtGet)
	//				stopArea2 = true
	//			}
	//
	//			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
	//				var (
	//					code int64
	//				)
	//
	//				code, err = uuc.uiRepo.UpdateUserRewardAreaTwo(ctx, tmpRecommendUser.ID, tmp, stopArea2)
	//				if code > 0 && err != nil {
	//					fmt.Println("错误分红社区：", err, tmpRecommendUser)
	//				}
	//
	//				if stopArea2 {
	//					stopUserIds[tmpRecommendUser.ID] = true // 出局
	//
	//					// 推荐人
	//					var (
	//						userRecommendArea *UserRecommend
	//					)
	//					if _, ok := userRecommendsMap[tmpRecommendUser.ID]; ok {
	//						userRecommendArea = userRecommendsMap[tmpRecommendUser.ID]
	//					} else {
	//						fmt.Println("错误分红社区，信息缺失7：", err, vBuys)
	//					}
	//
	//					if nil != userRecommendArea && "" != userRecommendArea.RecommendCode {
	//						var tmpRecommendAreaUserIds []string
	//						tmpRecommendAreaUserIds = strings.Split(userRecommendArea.RecommendCode, "D")
	//
	//						for _, vTmpRecommendAreaUserIds := range tmpRecommendAreaUserIds {
	//							if 0 >= len(vTmpRecommendAreaUserIds) {
	//								continue
	//							}
	//
	//							myUserRecommendAreaUserId, _ := strconv.ParseInt(vTmpRecommendAreaUserIds, 10, 64) // 最后一位是直推人
	//							if 0 >= myUserRecommendAreaUserId {
	//								continue
	//							}
	//
	//							// 减掉业绩
	//							err = uuc.uiRepo.UpdateUserMyTotalAmount(ctx, myUserRecommendAreaUserId, tmpRecommendUser.AmountUsdt)
	//							if err != nil {
	//								fmt.Println("错误分红社区：", err, vBuys)
	//							}
	//						}
	//					}
	//				}
	//
	//				return nil
	//			}); nil != err {
	//				fmt.Println("err reward daily area 2", err, vBuys)
	//			}
	//
	//			lastLevel = currentLevel
	//			lastLevelNum = tmpLastLevelNum
	//		}
	//	}
	//
	//}

	return nil, err
}

func (uuc *UserUseCase) AdminSetIspay(ctx context.Context, req *v1.AdminSetIspayRequest) (*v1.AdminSetIspayReply, error) {
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByAddressTwo(ctx, req.SendBody.Address)
	if nil != err {
		return nil, nil
	}

	if nil != user && 0 < user.ID {
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { //
			err = uuc.uiRepo.UpdateUserIspay(ctx, user.ID, req.SendBody.Amount)
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminSetBuyFour(ctx context.Context, req *v1.AdminSetIspayRequest) (*v1.AdminSetIspayReply, error) {
	// 推荐人
	var (
		err        error
		configs    []*Config
		priceOne   float64
		priceTwo   float64
		priceThree float64
		//recommendRate    float64
		//recommendRateTwo float64
	)

	var (
		user *User
	)
	user, err = uuc.repo.GetUserByAddressTwo(ctx, req.SendBody.Address)
	if nil != err {
		return nil, nil
	}

	// 配置
	configs, err = uuc.configRepo.GetConfigByKeys(ctx,
		"b_price_three_one",
		"b_price_three_two",
		"b_price_three_three",
		"recommend_reward_rate_one",
		"recommend_reward_rate_two",
	)
	if nil != err || nil == configs {
		return &v1.AdminSetIspayReply{
			Status: "稍后重试|err wait",
		}, nil
	}

	for _, vConfig := range configs {
		if "b_price_three_one" == vConfig.KeyName {
			priceOne, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "b_price_three_two" == vConfig.KeyName {
			priceTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		if "b_price_three_three" == vConfig.KeyName {
			priceThree, _ = strconv.ParseFloat(vConfig.Value, 10)
		}
		//if "recommend_reward_rate_one" == vConfig.KeyName {
		//	recommendRate, _ = strconv.ParseFloat(vConfig.Value, 10)
		//}
		//if "recommend_reward_rate_two" == vConfig.KeyName {
		//	recommendRateTwo, _ = strconv.ParseFloat(vConfig.Value, 10)
		//}
	}

	amount := req.SendBody.Amount
	amountIspay := float64(0)
	amountIspayPerDay := float64(0)
	price := float64(0)
	if 100 <= amount && 4000 >= amount {
		price = priceOne
		amountIspay = float64(amount) / price
		amountIspayPerDay = amountIspay / 300
	} else if 4500 <= amount && 10000 >= amount {
		price = priceTwo
		amountIspay = float64(amount) / price
		amountIspayPerDay = amountIspay / 600
	} else if 12000 <= amount && 100000 >= amount {
		price = priceThree
		amountIspay = float64(amount) / price
		amountIspayPerDay = amountIspay / 750
	} else {
		return &v1.AdminSetIspayReply{
			Status: "参数错误 |err amount",
		}, nil
	}

	if 0.000001 >= price || 0.000001 >= amountIspayPerDay || 0.000001 >= amountIspay {
		return &v1.AdminSetIspayReply{
			Status: "参数错误 |err amount",
		}, nil
	}

	//var (
	//	users    []*User
	//	usersMap map[int64]*User
	//)
	//users, err = uuc.ubRepo.GetAllUsersB(ctx)
	//if nil == users {
	//	return &v1.AdminSetIspayReply{
	//		Status: "参数错误 |err id",
	//	}, nil
	//}
	//
	//usersMap = make(map[int64]*User, 0)
	//for _, vUsers := range users {
	//	usersMap[vUsers.ID] = vUsers
	//}
	//
	//if _, ok := usersMap[user.ID]; !ok {
	//	fmt.Println("不存在用户")
	//	return &v1.AdminSetIspayReply{
	//		Status: "参数错误 |err id",
	//	}, nil
	//}
	//user = usersMap[user.ID]

	//var (
	//	amountRel   = float64(amount)
	//	userBalance *UserBalance
	//)
	//
	//userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
	//if nil == userBalance || nil != err {
	//	return &v1.BuyReply{
	//		Status: "参数错误 |err id",
	//	}, nil
	//}
	//
	//if amountRel > userBalance.BalanceUsdtFloat {
	//	return &v1.BuyReply{
	//		Status: "usdt余额不足|deposit usdt not enough",
	//	}, nil
	//}

	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserNewNewNewFour(ctx, user.ID, amount, amountIspay, amountIspayPerDay, "", "", "")
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资4", amount)
		return &v1.AdminSetIspayReply{
			Status: "参数错误 |err id",
		}, nil
	}

	return nil, nil
}

// AdminAddMoney  .
func (uuc *UserUseCase) AdminAddMoney(ctx context.Context, req *v1.AdminDailyAddMoneyRequest) (*v1.AdminDailyAddMoneyReply, error) {
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByAddressTwo(ctx, req.SendBody.Address)
	if nil != err {
		return nil, nil
	}

	if nil != user && 0 < user.ID {
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { //
			err = uuc.uiRepo.UpdateUserNewTwoNewThree(ctx, user.ID, uint64(req.SendBody.Usdt), user.Amount, "USDT")
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
	}

	return nil, nil
}

// AdminSubMoney  .
func (uuc *UserUseCase) AdminSubMoney(ctx context.Context, req *v1.AdminSubMoneyRequest) (*v1.AdminSubMoneyReply, error) {
	var (
		buyRecord *BuyRecord
		err       error
	)
	buyRecord, err = uuc.ubRepo.GetUserBuyById(req.Id)
	if nil != err {
		return &v1.AdminSubMoneyReply{}, err
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserSubBuyRecord(ctx, buyRecord.ID, buyRecord.UserId, buyRecord.Amount)
		if err != nil {
			fmt.Println("错误分红静态：", err, buyRecord)
			return err
		}

		return nil
	}); nil != err {
		fmt.Println("err sub daily", err, buyRecord)
		return &v1.AdminSubMoneyReply{}, err
	}

	//var (
	//	userRecommend *UserRecommend
	//)
	//userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, buyRecord.UserId)
	//if nil != err {
	//	return &v1.AdminSubMoneyReply{}, err
	//}
	//
	//if nil != userRecommend && "" != userRecommend.RecommendCode {
	//	var tmpRecommendUserIds []string
	//	tmpRecommendUserIds = strings.Split(userRecommend.RecommendCode, "D")
	//	for j := len(tmpRecommendUserIds) - 1; j >= 0; j-- {
	//		if 0 >= len(tmpRecommendUserIds[j]) {
	//			continue
	//		}
	//
	//		myUserRecommendUserId, _ := strconv.ParseInt(tmpRecommendUserIds[j], 10, 64) // 最后一位是直推人
	//		if 0 >= myUserRecommendUserId {
	//			continue
	//		}
	//		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
	//			// 减掉业绩
	//			err = uuc.uiRepo.UpdateUserMyTotalAmountSub(ctx, myUserRecommendUserId, buyRecord.Amount)
	//			if err != nil {
	//				fmt.Println("错误sub：", err, buyRecord, myUserRecommendUserId)
	//				return err
	//			}
	//
	//			return nil
	//		}); nil != err {
	//			fmt.Println("err sub 业绩更新", err, buyRecord)
	//			continue
	//		}
	//	}
	//}

	return nil, nil
}

func (uuc *UserUseCase) AdminAddMoneyThree(ctx context.Context, req *v1.AdminDailyAddMoneyTwoRequest) (*v1.AdminDailyAddMoneyTwoReply, error) {
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByAddressTwo(ctx, req.SendBody.Address)
	if nil != err {
		return nil, nil
	}

	amount := req.SendBody.Usdt
	four := 55
	if 100 == amount {
		amount = 100
	} else if 300 == amount {
		amount = 300
		four = 63
	} else if 500 == amount {
		amount = 500
		four = 64
	} else if 1000 == amount {
		amount = 1000
		four = 65
	} else if 5000 == amount {
		amount = 5000
		four = 66
	} else if 10000 == amount {
		amount = 10000
		four = 67
	} else if 15000 == amount {
		amount = 15000
		four = 68
	} else if 30000 == amount {
		amount = 30000
		four = 69
	} else if 50000 == amount {
		amount = 50000
		four = 70
	} else if 100000 == amount {
		amount = 100000
		four = 71
	} else if 150000 == amount {
		amount = 150000
		four = 72
	} else {
		return &v1.AdminDailyAddMoneyTwoReply{}, nil
	}

	one := ""
	two := ""
	three := ""
	if "1" != user.One {
		one += user.One
	}
	if "1" != user.Two {
		one += user.Two
	}
	if "1" != user.Three {
		one += user.Three
	}
	if "1" != user.Four {
		one += user.Four
	}
	if "1" != user.Five {
		one += user.Five
	}
	if "1" != user.Six {
		two = user.Six
	}
	if "1" != user.Seven {
		three = user.Seven
	}

	// 入金
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.uiRepo.UpdateUserNewTwoNewTwoNew(ctx, user.ID, uint64(amount), one, two, three, int64(four))
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println(err, "错误投资3", amount)
	}

	return nil, nil
}

// AdminAddMoneyTwo  .
func (uuc *UserUseCase) AdminAddMoneyTwo(ctx context.Context, req *v1.AdminDailyAddMoneyTwoRequest) (*v1.AdminDailyAddMoneyTwoReply, error) {
	var (
		user *User
		err  error
	)
	user, err = uuc.repo.GetUserByAddressTwo(ctx, req.SendBody.Address)
	if nil != err {
		return nil, nil
	}

	if nil != user && 0 < user.ID {
		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { //
			err = uuc.uiRepo.UpdateUserUsdtFloat(ctx, user.ID, float64(req.SendBody.Usdt), 0, "USDT")
			if nil != err {
				return err
			}

			return nil
		}); nil != err {
			return nil, err
		}
	}

	return nil, nil
}

// AdminRecommendLevelUpdate  .
func (uuc *UserUseCase) AdminRecommendLevelUpdate(ctx context.Context, req *v1.AdminRecommendLevelRequest) (*v1.AdminRecommendLevelReply, error) {
	var (
		err error
	)

	err = uuc.uiRepo.UpdateUserRecommendLevel(ctx, req.SendBody.UserId, uint64(req.SendBody.Level))
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func (uuc *UserUseCase) AdminDailyAreaReward(ctx context.Context, req *v1.AdminDailyLocationRewardRequest) (*v1.AdminDailyLocationRewardReply, error) {
	var (
		userLocations []*LocationNew
		configs       []*Config
		bPrice        int64
		bPriceBase    int64
		areaOne       int64
		areaTwo       int64
		areaThree     int64
		areaFour      int64
		areaFive      int64
		areaNumOne    int64
		areaNumTwo    int64
		areaNumThree  int64
		areaNumFour   int64
		areaNumFive   int64
		one           int64
		two           int64
		three         int64
		four          int64
		total         int64
		feeRate       int64
		err           error
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx,
		"location_reward_rate", "b_price", "b_price_base", "exchange_rate",
		"recommend_one_rate", "recommend_two_rate",
		"recommend_three_rate", "recommend_four_rate",
		"recommend_five_rate", "recommend_six_rate",
		"recommend_seven_rate", "recommend_eight_rate",
		"area_one", "area_two", "area_three", "area_four", "area_five",
		"area_num_one", "area_num_two", "area_num_three", "area_num_four", "area_num_five", "one", "two", "three", "four", "total",
	)
	if nil != configs {
		for _, vConfig := range configs {
			if "b_price" == vConfig.KeyName {
				bPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "b_price_base" == vConfig.KeyName {
				bPriceBase, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_one" == vConfig.KeyName {
				areaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_two" == vConfig.KeyName {
				areaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_three" == vConfig.KeyName {
				areaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_four" == vConfig.KeyName {
				areaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_five" == vConfig.KeyName {
				areaFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_one" == vConfig.KeyName {
				areaNumOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_two" == vConfig.KeyName {
				areaNumTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_three" == vConfig.KeyName {
				areaNumThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_four" == vConfig.KeyName {
				areaNumFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "area_num_five" == vConfig.KeyName {
				areaNumFive, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "one" == vConfig.KeyName {
				one, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "two" == vConfig.KeyName {
				two, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "three" == vConfig.KeyName {
				three, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "four" == vConfig.KeyName {
				four, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "total" == vConfig.KeyName {
				total, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "exchange_rate" == vConfig.KeyName {
				feeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	var (
		users    []*User
		usersMap map[int64]*User
	)
	users, err = uuc.repo.GetAllUsers(ctx)
	if nil == users {
		fmt.Println("今日分红错误用户获取失败")
		return nil, nil
	}

	usersMap = make(map[int64]*User, 0)
	for _, vUsers := range users {
		usersMap[vUsers.ID] = vUsers
	}

	// 获取今日收益
	var (
		day               = -1
		userLocationsYes  []*LocationNew
		rewardLocationYes int64
	)
	// 全网
	userLocationsYes, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day)
	for _, userLocationYes := range userLocationsYes {
		rewardLocationYes += userLocationYes.Usdt
	}

	if 0 >= rewardLocationYes {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}

	// 团队奖励
	userLocationsOne := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		if _, ok := usersMap[vUserLocations.UserId]; ok {
			if 1 == usersMap[vUserLocations.UserId].Lock {
				continue
			}
		}

		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaOne <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 1 <= vUserLocations.LastLevel {
				userLocationsOne = append(userLocationsOne, vUserLocations)
			}

		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaOne <= vUserLocations.Total+vUserLocations.TotalThree || 1 <= vUserLocations.LastLevel {
				userLocationsOne = append(userLocationsOne, vUserLocations)
			}

		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaOne <= vUserLocations.TotalTwo+vUserLocations.Total || 1 <= vUserLocations.LastLevel {
				userLocationsOne = append(userLocationsOne, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsOne) {
		rewardLocationYesOne := rewardLocationYes / 1000 * areaNumOne / int64(len(userLocationsOne))
		if 0 < rewardLocationYesOne {
			for _, vUserLocationsItem := range userLocationsOne {
				// 奖励
				tmpCurrentReward := rewardLocationYesOne
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.UserId, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate, vUserLocationsItem.CurrentMax) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 1, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
								if nil != err {
									return err
								}

								if nil != currentLocation && 0 < currentLocation.Top {
									tmpTop = currentLocation.Top
									tmpTopNum = currentLocation.TopNum
									continue
								}

								break
							}
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	// 团队奖励
	userLocationsTwo := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		if _, ok := usersMap[vUserLocations.UserId]; ok {
			if 1 == usersMap[vUserLocations.UserId].Lock {
				continue
			}
		}

		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaTwo <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 2 <= vUserLocations.LastLevel {
				userLocationsTwo = append(userLocationsTwo, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaTwo <= vUserLocations.Total+vUserLocations.TotalThree || 2 <= vUserLocations.LastLevel {
				userLocationsTwo = append(userLocationsTwo, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaTwo <= vUserLocations.TotalTwo+vUserLocations.Total || 2 <= vUserLocations.LastLevel {
				userLocationsTwo = append(userLocationsTwo, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsTwo) {
		rewardLocationYesTwo := rewardLocationYes / 1000 * areaNumTwo / int64(len(userLocationsTwo))
		if 0 < rewardLocationYesTwo {
			for _, vUserLocationsItem := range userLocationsTwo {
				// 奖励
				tmpCurrentReward := rewardLocationYesTwo
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.UserId, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate, vUserLocationsItem.CurrentMax) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 2, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
								if nil != err {
									return err
								}

								if nil != currentLocation && 0 < currentLocation.Top {
									tmpTop = currentLocation.Top
									tmpTopNum = currentLocation.TopNum
									continue
								}

								break
							}
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	userLocationsThree := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		if _, ok := usersMap[vUserLocations.UserId]; ok {
			if 1 == usersMap[vUserLocations.UserId].Lock {
				continue
			}
		}

		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaThree <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 3 <= vUserLocations.LastLevel {
				userLocationsThree = append(userLocationsThree, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaThree <= vUserLocations.Total+vUserLocations.TotalThree || 3 <= vUserLocations.LastLevel {
				userLocationsThree = append(userLocationsThree, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaThree <= vUserLocations.TotalTwo+vUserLocations.Total || 3 <= vUserLocations.LastLevel {
				userLocationsThree = append(userLocationsThree, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsThree) {
		rewardLocationYesThree := rewardLocationYes / 1000 * areaNumThree / int64(len(userLocationsThree))
		if 0 < rewardLocationYesThree {
			for _, vUserLocationsItem := range userLocationsThree {
				// 奖励
				tmpCurrentReward := rewardLocationYesThree
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.UserId, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate, vUserLocationsItem.CurrentMax) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 3, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
								if nil != err {
									return err
								}

								if nil != currentLocation && 0 < currentLocation.Top {
									tmpTop = currentLocation.Top
									tmpTopNum = currentLocation.TopNum
									continue
								}

								break
							}
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	userLocationsFour := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		if _, ok := usersMap[vUserLocations.UserId]; ok {
			if 1 == usersMap[vUserLocations.UserId].Lock {
				continue
			}
		}

		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaFour <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 4 <= vUserLocations.LastLevel {
				userLocationsFour = append(userLocationsFour, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaFour <= vUserLocations.Total+vUserLocations.TotalThree || 4 <= vUserLocations.LastLevel {
				userLocationsFour = append(userLocationsFour, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaFour <= vUserLocations.TotalTwo+vUserLocations.Total || 4 <= vUserLocations.LastLevel {
				userLocationsFour = append(userLocationsFour, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsFour) {
		rewardLocationYesFour := rewardLocationYes / 1000 * areaNumFour / int64(len(userLocationsFour))
		if 0 < rewardLocationYesFour {
			for _, vUserLocationsItem := range userLocationsFour {
				// 奖励
				tmpCurrentReward := rewardLocationYesFour
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.UserId, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate, vUserLocationsItem.CurrentMax) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 4, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
								if nil != err {
									return err
								}

								if nil != currentLocation && 0 < currentLocation.Top {
									tmpTop = currentLocation.Top
									tmpTopNum = currentLocation.TopNum
									continue
								}

								break
							}
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	userLocationsFive := make([]*LocationNew, 0)
	userLocations, err = uuc.locationRepo.GetRunningLocations(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}
	for _, vUserLocations := range userLocations {
		if _, ok := usersMap[vUserLocations.UserId]; ok {
			if 1 == usersMap[vUserLocations.UserId].Lock {
				continue
			}
		}

		// 1大区
		if vUserLocations.Total >= vUserLocations.TotalTwo && vUserLocations.Total >= vUserLocations.TotalThree {
			if areaFive <= vUserLocations.TotalTwo+vUserLocations.TotalThree || 5 <= vUserLocations.LastLevel {
				userLocationsFive = append(userLocationsFive, vUserLocations)
			}
		} else if vUserLocations.TotalTwo >= vUserLocations.Total && vUserLocations.TotalTwo >= vUserLocations.TotalThree {
			if areaFive <= vUserLocations.Total+vUserLocations.TotalThree || 5 <= vUserLocations.LastLevel {
				userLocationsFive = append(userLocationsFive, vUserLocations)
			}
		} else if vUserLocations.TotalThree >= vUserLocations.Total && vUserLocations.TotalThree >= vUserLocations.TotalTwo {
			if areaFive <= vUserLocations.TotalTwo+vUserLocations.Total || 5 <= vUserLocations.LastLevel {
				userLocationsFive = append(userLocationsFive, vUserLocations)
			}
		}
	}

	if 0 < len(userLocationsFive) {
		rewardLocationYesFive := rewardLocationYes / 1000 * areaNumFive / int64(len(userLocationsFive))
		if 0 < rewardLocationYesFive {
			for _, vUserLocationsItem := range userLocationsFive {
				// 奖励
				tmpCurrentReward := rewardLocationYesFive
				bLocationRewardAmount := tmpCurrentReward * bPriceBase / bPrice

				if 0 < tmpCurrentReward {
					if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if vUserLocationsItem.Current+tmpCurrentReward >= vUserLocationsItem.CurrentMax { // 占位分红人分满停止
							vUserLocationsItem.Status = "stop"
							vUserLocationsItem.StopDate = time.Now().UTC().Add(8 * time.Hour)

							tmpCurrentReward = vUserLocationsItem.CurrentMax - vUserLocationsItem.Current
							bLocationRewardAmount = tmpCurrentReward * bPriceBase / bPrice
						}

						var tmpMaxNew int64
						if vUserLocationsItem.CurrentMaxNew < vUserLocationsItem.CurrentMax {
							tmpMaxNew = vUserLocationsItem.CurrentMax - vUserLocationsItem.CurrentMaxNew
						}

						if 0 < tmpCurrentReward {
							err = uuc.locationRepo.UpdateLocationNewNew(ctx, vUserLocationsItem.ID, vUserLocationsItem.UserId, vUserLocationsItem.Status, tmpCurrentReward, tmpMaxNew, bLocationRewardAmount, vUserLocationsItem.StopDate, vUserLocationsItem.CurrentMax) // 分红占位数据修改
							if nil != err {
								return err
							}

							_, err = uuc.ubRepo.AreaRewardBiw(ctx, vUserLocationsItem.UserId, bLocationRewardAmount, tmpCurrentReward, 5, vUserLocationsItem.Status, tmpMaxNew, feeRate)
							if nil != err {
								return err
							}
						}

						// 业绩减掉
						if "stop" == vUserLocationsItem.Status {
							tmpTop := vUserLocationsItem.Top
							tmpTopNum := vUserLocationsItem.TopNum
							for j := 0; j < 10000 && 0 < tmpTop && 0 < tmpTopNum; j++ {
								err = uuc.locationRepo.UpdateLocationNewTotalSub(ctx, tmpTop, tmpTopNum, vUserLocationsItem.Usdt/100000)
								if nil != err {
									return err
								}

								var (
									currentLocation *LocationNew
								)
								currentLocation, err = uuc.locationRepo.GetLocationById(ctx, tmpTop)
								if nil != err {
									return err
								}

								if nil != currentLocation && 0 < currentLocation.Top {
									tmpTop = currentLocation.Top
									tmpTopNum = currentLocation.TopNum
									continue
								}

								break
							}
						}

						return nil
					}); nil != err {
						fmt.Println("err reward daily three", err, vUserLocationsItem)
						continue
					}
				}
			}
		}
	}

	// 全网前天
	var (
		rewardFourYes *Reward
	)
	rewardLocationYes = rewardLocationYes / 100 * total
	fmt.Println("今天：", rewardLocationYes)
	rewardFourYes, err = uuc.ubRepo.GetRewardFourYes(ctx) // 推荐人奖励
	if nil == err && nil != rewardFourYes {
		rewardLocationYes += rewardFourYes.Amount
	}
	fmt.Println("今天+昨日沉淀：", rewardLocationYes)
	// 全球
	//totalReward := rewardLocationYes/100/100*70*total + rewardLocationBef/100/100*30*total
	totalReward := rewardLocationYes / 100 * 70

	var (
		fourUserRecommendTotal map[int64]int64
	)

	fourUserRecommendTotal = make(map[int64]int64, 0)
	for _, userLocationYes := range userLocationsYes {
		// 获取直推

		var (
			fourUserRecommend         *UserRecommend
			myFourUserRecommendUserId int64
			//myFourRecommendUser *User
		)
		fourUserRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, userLocationYes.UserId)
		if nil == fourUserRecommend {
			continue
		}

		if "" != fourUserRecommend.RecommendCode {
			tmpFourRecommendUserIds := strings.Split(fourUserRecommend.RecommendCode, "D")
			if 2 <= len(tmpFourRecommendUserIds) {
				myFourUserRecommendUserId, _ = strconv.ParseInt(tmpFourRecommendUserIds[len(tmpFourRecommendUserIds)-1], 10, 64) // 最后一位是直推人
			}
			//myFourRecommendUser, err = uuc.repo.GetUserById(ctx, myFourUserRecommendUserId)
			//if nil != err {
			//	return nil, err
			//}

			if _, ok := fourUserRecommendTotal[myFourUserRecommendUserId]; ok {
				fourUserRecommendTotal[myFourUserRecommendUserId] += userLocationYes.Usdt
			} else {
				fourUserRecommendTotal[myFourUserRecommendUserId] = userLocationYes.Usdt
			}
		}
	}

	if 0 >= len(fourUserRecommendTotal) {
		return &v1.AdminDailyLocationRewardReply{}, nil
	}

	// 前四名
	type KeyValuePair struct {
		Key   int64
		Value int64
	}
	var keyValuePairs []KeyValuePair
	for key, value := range fourUserRecommendTotal {
		keyValuePairs = append(keyValuePairs, KeyValuePair{key, value})
	}

	// 按值排序切片
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// 获取前四项
	var topFour []KeyValuePair
	if 4 <= len(keyValuePairs) {
		topFour = keyValuePairs[:4]
	} else {
		topFour = keyValuePairs[:len(keyValuePairs)]
	}

	for k, vTopFour := range topFour {
		var (
			tmpMyRecommendAmount int64
		)
		if 0 == k {
			tmpMyRecommendAmount = totalReward / 100 * one
		} else if 1 == k {
			tmpMyRecommendAmount = totalReward / 100 * two
		} else if 2 == k {
			tmpMyRecommendAmount = totalReward / 100 * three
		} else if 3 == k {
			tmpMyRecommendAmount = totalReward / 100 * four
		}

		if 0 >= tmpMyRecommendAmount {
			continue
		}

		if 0 < tmpMyRecommendAmount {
			if _, ok := usersMap[vTopFour.Key]; ok {
				if 1 == usersMap[vTopFour.Key].Lock {
					continue
				}
			}

			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				_, err = uuc.ubRepo.FourRewardBiw(ctx, vTopFour.Key, tmpMyRecommendAmount, int64(k+1)) // 推荐人奖励
				if nil != err {
					return err
				}

				return nil
			}); nil != err {
				fmt.Println("err reward daily four", err, vTopFour)
				continue
			}
		}
	}

	fmt.Println("今日沉淀", rewardLocationYes/100*30)
	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = uuc.ubRepo.FourRewardYes(ctx, rewardLocationYes/100*30) // 推荐人奖励
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		fmt.Println("err reward daily four yes", err, rewardLocationYes/100*30)
	}

	return &v1.AdminDailyLocationRewardReply{}, nil
}

func (uuc *UserUseCase) AdminUpdateLocationNewMax(ctx context.Context, req *v1.AdminUpdateLocationNewMaxRequest) (*v1.AdminUpdateLocationNewMaxReply, error) {
	var (
		err error
	)
	res := &v1.AdminUpdateLocationNewMaxReply{}
	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloat *= 100000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloat, 'f', -1, 64), 10, 64)

	_, err = uuc.ubRepo.UpdateLocationNewMax(ctx, req.SendBody.UserId, amount)

	if nil != err {
		return res, err
	}

	return nil, err
}

func (uuc *UserUseCase) AdminDailyLocationRewardNew(ctx context.Context, req *v1.AdminDailyLocationRewardNewRequest) (*v1.AdminDailyLocationRewardNewReply, error) {
	var (
		userLocations    []*LocationNew
		userLocationsMap map[int64]*LocationNew
		userLocations1   []*LocationNew
		userLocations2   []*LocationNew
		userLocations3   []*LocationNew
		v1r              int64
		v2r              int64
		v3r              int64
		configs          []*Config
		amount           int64
		amountV1         int64
		amountV2         int64
		amountV3         int64
		v1Count          int64
		v2Count          int64
		v3Count          int64
		err              error
		userInfos        []*UserInfo
	)

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx,
		"v1", "v2", "v3",
	)

	if nil != configs {
		for _, vConfig := range configs {
			if "v1" == vConfig.KeyName {
				v1r, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "v2" == vConfig.KeyName {
				v2r, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "v3" == vConfig.KeyName {
				v3r, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	amount, err = uuc.ubRepo.GetSystemWithdrawUsdtFeeTotalToday(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardNewReply{}, nil
	}

	// 获取手动设置的
	userInfos, err = uuc.uiRepo.GetUserInfosByVipAndLockVip(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardNewReply{}, nil
	}

	if nil != userInfos {
		for _, v := range userInfos {
			if 3 == v.Vip {
				v3Count += 1
			}

			if 2 == v.Vip {
				v2Count += 1
			}

			if 1 == v.Vip {
				v1Count += 1
			}
		}
	}

	userLocations, err = uuc.locationRepo.GetAllLocationsNew2(ctx)
	if nil != err {
		return &v1.AdminDailyLocationRewardNewReply{}, nil
	}

	userLocationsMap = make(map[int64]*LocationNew, 0)
	userLocations3 = make([]*LocationNew, 0)
	userLocations2 = make([]*LocationNew, 0)
	userLocations1 = make([]*LocationNew, 0)
	for _, vUserLocations := range userLocations {

		if _, ok := userLocationsMap[vUserLocations.UserId]; ok {
			continue
		}

		var (
			userInfo *UserInfo
		)

		userLocationsMap[vUserLocations.UserId] = vUserLocations

		userInfo, err = uuc.uiRepo.GetUserInfoByUserId(ctx, vUserLocations.UserId)
		if nil != err {
			continue
		}

		if 3 == userInfo.Vip {
			v3Count += 1
			userLocations3 = append(userLocations3, vUserLocations)
		}

		if 2 == userInfo.Vip {
			v2Count += 1
			userLocations2 = append(userLocations2, vUserLocations)
		}

		if 1 == userInfo.Vip {
			v1Count += 1
			userLocations1 = append(userLocations1, vUserLocations)
		}
	}

	if v1Count > 0 {
		amountV1 = amount * v1r / 100 / v1Count
	}

	if v2Count > 0 {
		amountV2 = amount * v2r / 100 / v2Count
	}

	if v3Count > 0 {
		amountV3 = amount * v3r / 100 / v3Count
	}

	if nil != userInfos {
		for _, v := range userInfos {
			var tmpAmount int64
			if 3 == v.Vip {
				tmpAmount = amountV3
			}

			if 2 == v.Vip {
				tmpAmount = amountV2
			}

			if 1 == v.Vip {
				tmpAmount = amountV1
			}
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = uuc.ubRepo.NormalReward4(ctx, v.UserId, tmpAmount, 0)
				if nil != err {
					return err
				}
				return nil
			}); nil != err {
				continue
			}
		}
	}

	for _, vUserLocations1 := range userLocations1 {

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpStatus := vUserLocations1.Status // 现在还在运行中

			// 奖励usdt
			tmpRewardAmount := amountV1

			vUserLocations1.Status = "running"
			vUserLocations1.Current += tmpRewardAmount

			tmpRewardAmount2 := tmpRewardAmount
			if vUserLocations1.Current >= vUserLocations1.CurrentMax { // 占位分红人分满停止
				vUserLocations1.Status = "stop"
				if "running" == tmpStatus {
					vUserLocations1.StopDate = time.Now().UTC().Add(8 * time.Hour)
					tmpRewardAmount2 = tmpRewardAmount - (vUserLocations1.Current - vUserLocations1.CurrentMax)
				} else {
					tmpRewardAmount2 = 0
				}
			}

			if 0 < tmpRewardAmount {
				err = uuc.locationRepo.UpdateLocationNew2(ctx, vUserLocations1.ID, vUserLocations1.Status, tmpRewardAmount, vUserLocations1.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}
				_, err = uuc.ubRepo.NormalReward3(ctx, vUserLocations1.UserId, tmpRewardAmount, tmpRewardAmount2, vUserLocations1.ID, tmpStatus, vUserLocations1.Status) // 直推人奖励
				if nil != err {
					return err
				}
			}
			return nil
		}); nil != err {
			continue
		}
	}

	for _, vUserLocations3 := range userLocations3 {

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpStatus := vUserLocations3.Status // 现在还在运行中

			// 奖励usdt
			tmpRewardAmount := amountV3

			vUserLocations3.Status = "running"
			vUserLocations3.Current += tmpRewardAmount

			tmpRewardAmount2 := tmpRewardAmount
			if vUserLocations3.Current >= vUserLocations3.CurrentMax { // 占位分红人分满停止
				vUserLocations3.Status = "stop"
				if "running" == tmpStatus {
					vUserLocations3.StopDate = time.Now().UTC().Add(8 * time.Hour)
					tmpRewardAmount2 = tmpRewardAmount - (vUserLocations3.Current - vUserLocations3.CurrentMax)
				} else {
					tmpRewardAmount2 = 0
				}
			}

			if 0 < tmpRewardAmount {
				err = uuc.locationRepo.UpdateLocationNew2(ctx, vUserLocations3.ID, vUserLocations3.Status, tmpRewardAmount, vUserLocations3.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}
				_, err = uuc.ubRepo.NormalReward3(ctx, vUserLocations3.UserId, tmpRewardAmount, tmpRewardAmount2, vUserLocations3.ID, tmpStatus, vUserLocations3.Status) // 直推人奖励
				if nil != err {
					return err
				}
			}
			return nil
		}); nil != err {
			continue
		}
	}

	for _, vUserLocations2 := range userLocations2 {

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			tmpStatus := vUserLocations2.Status // 现在还在运行中

			// 奖励usdt
			tmpRewardAmount := amountV2

			vUserLocations2.Status = "running"
			vUserLocations2.Current += tmpRewardAmount

			tmpRewardAmount2 := tmpRewardAmount
			if vUserLocations2.Current >= vUserLocations2.CurrentMax { // 占位分红人分满停止
				vUserLocations2.Status = "stop"
				if "running" == tmpStatus {
					vUserLocations2.StopDate = time.Now().UTC().Add(8 * time.Hour)
					tmpRewardAmount2 = tmpRewardAmount - (vUserLocations2.Current - vUserLocations2.CurrentMax)
				} else {
					tmpRewardAmount2 = 0
				}
			}

			if 0 < tmpRewardAmount {
				err = uuc.locationRepo.UpdateLocationNew2(ctx, vUserLocations2.ID, vUserLocations2.Status, tmpRewardAmount, vUserLocations2.StopDate) // 分红占位数据修改
				if nil != err {
					return err
				}
				_, err = uuc.ubRepo.NormalReward3(ctx, vUserLocations2.UserId, tmpRewardAmount, tmpRewardAmount2, vUserLocations2.ID, tmpStatus, vUserLocations2.Status) // 直推人奖励
				if nil != err {
					return err
				}
			}
			return nil
		}); nil != err {
			continue
		}
	}

	return &v1.AdminDailyLocationRewardNewReply{}, nil
}

func (uuc *UserUseCase) AdminDailyRecommendReward(ctx context.Context, req *v1.AdminDailyRecommendRewardRequest) (*v1.AdminDailyRecommendRewardReply, error) {

	var (
		users                  []*User
		userLocations          []*LocationNew
		configs                []*Config
		recommendAreaOne       int64
		recommendAreaOneRate   int64
		recommendAreaTwo       int64
		recommendAreaTwoRate   int64
		recommendAreaThree     int64
		recommendAreaThreeRate int64
		recommendAreaFour      int64
		recommendAreaFourRate  int64
		fee                    int64
		rewardRate             int64
		coinPrice              int64
		coinRewardRate         int64
		day                    = -1
		err                    error
	)

	if 1 == req.Day {
		day = 0
	}

	// 全网手续费
	userLocations, err = uuc.locationRepo.GetLocationDailyYesterday(ctx, day)
	if nil != err {
		return nil, err
	}
	for _, userLocation := range userLocations {
		fee += userLocation.CurrentMax * 100 / userLocation.OutRate
	}
	if 0 >= fee {
		return &v1.AdminDailyRecommendRewardReply{}, nil
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "recommend_area_one",
		"recommend_area_one_rate", "recommend_area_two_rate", "recommend_area_three_rate", "recommend_area_four_rate",
		"recommend_area_two", "recommend_area_three", "recommend_area_four", "coin_price", "coin_reward_rate", "reward_rate")
	if nil != configs {
		for _, vConfig := range configs {
			if "recommend_area_one" == vConfig.KeyName {
				recommendAreaOne, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_one_rate" == vConfig.KeyName {
				recommendAreaOneRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_two" == vConfig.KeyName {
				recommendAreaTwo, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_two_rate" == vConfig.KeyName {
				recommendAreaTwoRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_three" == vConfig.KeyName {
				recommendAreaThree, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_three_rate" == vConfig.KeyName {
				recommendAreaThreeRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_four" == vConfig.KeyName {
				recommendAreaFour, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "recommend_area_four_rate" == vConfig.KeyName {
				recommendAreaFourRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_price" == vConfig.KeyName {
				coinPrice, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "coin_reward_rate" == vConfig.KeyName {
				coinRewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "reward_rate" == vConfig.KeyName {
				rewardRate, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	users, err = uuc.repo.GetAllUsers(ctx)
	if nil != err {
		return nil, err
	}

	level1 := make(map[int64]int64, 0)
	level2 := make(map[int64]int64, 0)
	level3 := make(map[int64]int64, 0)
	level4 := make(map[int64]int64, 0)

	for _, user := range users {
		var userArea *UserArea
		userArea, err = uuc.urRepo.GetUserArea(ctx, user.ID)
		if nil != err {
			continue
		}

		if userArea.Level > 0 {
			if userArea.Level >= 1 {
				level1[user.ID] = user.ID
			}
			if userArea.Level >= 2 {
				level2[user.ID] = user.ID
			}
			if userArea.Level >= 3 {
				level3[user.ID] = user.ID
			}
			if userArea.Level >= 4 {
				level4[user.ID] = user.ID
			}
			continue
		}

		var userRecommend *UserRecommend
		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
		if nil != err {
			continue
		}

		// 伞下业绩
		var (
			myRecommendUsers   []*UserRecommend
			userAreas          []*UserArea
			maxAreaAmount      int64
			areaAmount         int64
			myRecommendUserIds []int64
		)
		myCode := userRecommend.RecommendCode + "D" + strconv.FormatInt(user.ID, 10)
		myRecommendUsers, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
		if nil == err {
			// 找直推
			for _, vMyRecommendUsers := range myRecommendUsers {
				myRecommendUserIds = append(myRecommendUserIds, vMyRecommendUsers.UserId)
			}
		}
		if 0 < len(myRecommendUserIds) {
			userAreas, err = uuc.urRepo.GetUserAreas(ctx, myRecommendUserIds)
			if nil == err {
				var (
					tmpTotalAreaAmount int64
				)
				for _, vUserAreas := range userAreas {
					tmpAreaAmount := vUserAreas.Amount + vUserAreas.SelfAmount
					tmpTotalAreaAmount += tmpAreaAmount
					if tmpAreaAmount > maxAreaAmount {
						maxAreaAmount = tmpAreaAmount
					}
				}

				areaAmount = tmpTotalAreaAmount - maxAreaAmount
			}
		}

		// 比较级别
		if areaAmount >= recommendAreaOne*100000 {
			level1[user.ID] = user.ID
		}

		if areaAmount >= recommendAreaTwo*100000 {
			level2[user.ID] = user.ID
		}

		if areaAmount >= recommendAreaThree*100000 {
			level3[user.ID] = user.ID
		}

		if areaAmount >= recommendAreaFour*100000 {
			level4[user.ID] = user.ID
		}
	}
	fmt.Println(level4, level3, level2, level1)
	// 分红
	fee /= 100000 // 这里多除五个0
	fmt.Println(fee)
	if 0 < len(level1) {
		feeLevel1 := fee * recommendAreaOneRate / 100 / int64(len(level1))
		feeLevel1 *= 100000

		for _, vLevel1 := range level1 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel1)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel1Usdt := feeLevel1 * rewardRate / 100
				feeLevel1Coin := feeLevel1 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel1
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel1 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel1Usdt = tmpLastAmount * rewardRate / 100
						feeLevel1Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel1 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel1, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel1, feeLevel1, feeLevel1Usdt, feeLevel1Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	// 分红
	if 0 < len(level2) {
		feeLevel2 := fee * recommendAreaTwoRate / 100 / int64(len(level2))
		feeLevel2 *= 100000
		for _, vLevel2 := range level2 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel2)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel2Usdt := feeLevel2 * rewardRate / 100
				feeLevel2Coin := feeLevel2 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel2
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel2 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel2Usdt = tmpLastAmount * rewardRate / 100
						feeLevel2Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel2 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel2, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel2, feeLevel2, feeLevel2Usdt, feeLevel2Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	// 分红
	if 0 < len(level3) {
		feeLevel3 := fee * recommendAreaThreeRate / 100 / int64(len(level3))
		feeLevel3 *= 100000
		for _, vLevel3 := range level3 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel3)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel3Usdt := feeLevel3 * rewardRate / 100
				feeLevel3Coin := feeLevel3 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel3
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel3 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel3Usdt = tmpLastAmount * rewardRate / 100
						feeLevel3Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice
					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel3 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel3, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel3, feeLevel3, feeLevel3Usdt, feeLevel3Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	// 分红
	if 0 < len(level4) {
		feeLevel4 := fee * recommendAreaFourRate / 100 / int64(len(level4))
		feeLevel4 *= 100000
		for _, vLevel4 := range level4 {
			if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				var myLocationLast *LocationNew
				// 获取当前用户的占位信息，已经有运行中的跳过
				myLocationLast, err = uuc.locationRepo.GetMyLocationLast(ctx, vLevel4)
				if nil == myLocationLast { // 无占位信息
					return err
				}

				feeLevel4Usdt := feeLevel4 * rewardRate / 100
				feeLevel4Coin := feeLevel4 * coinRewardRate / 100 * 1000 / coinPrice

				tmpCurrentStatus := myLocationLast.Status // 现在还在运行中
				myLocationLast.Status = "running"
				myLocationLast.Current += feeLevel4
				if myLocationLast.Current >= myLocationLast.CurrentMax { // 占位分红人分满停止
					if "running" == tmpCurrentStatus {
						myLocationLast.StopDate = time.Now().UTC().Add(8 * time.Hour)

						tmpLastAmount := feeLevel4 - (myLocationLast.Current - myLocationLast.CurrentMax)
						feeLevel4Usdt = tmpLastAmount * rewardRate / 100
						feeLevel4Coin = tmpLastAmount * coinRewardRate / 100 * 1000 / coinPrice

					}
					myLocationLast.Status = "stop"
				}

				if 0 < feeLevel4 {
					err = uuc.locationRepo.UpdateLocationNew(ctx, myLocationLast.ID, myLocationLast.Status, feeLevel4, myLocationLast.StopDate) // 分红占位数据修改
					if nil != err {
						return err
					}

					_, err = uuc.ubRepo.UserDailyRecommendArea(ctx, vLevel4, feeLevel4, feeLevel4Usdt, feeLevel4Coin, tmpCurrentStatus)
					if nil != err {
						return err
					}

				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	return &v1.AdminDailyRecommendRewardReply{}, nil
}

func (uuc *UserUseCase) CheckAndInsertRecommendArea(ctx context.Context, req *v1.CheckAndInsertRecommendAreaRequest) (*v1.CheckAndInsertRecommendAreaReply, error) {

	var (
		userRecommends         []*UserRecommend
		userRecommendAreaCodes []string
		userRecommendAreas     []*UserRecommendArea
		err                    error
	)
	userRecommends, err = uuc.urRepo.GetUserRecommends(ctx)
	if nil != err {
		return &v1.CheckAndInsertRecommendAreaReply{}, nil
	}

	for _, vUserRecommends := range userRecommends {
		tmp := vUserRecommends.RecommendCode + "D" + strconv.FormatInt(vUserRecommends.UserId, 10)
		tmpNoHas := true
		for k, vUserRecommendAreaCodes := range userRecommendAreaCodes {
			if strings.HasPrefix(vUserRecommendAreaCodes, tmp) {
				tmpNoHas = false
			} else if strings.HasPrefix(tmp, vUserRecommendAreaCodes) {
				userRecommendAreaCodes[k] = tmp
				tmpNoHas = false
			}
		}

		if tmpNoHas {
			userRecommendAreaCodes = append(userRecommendAreaCodes, tmp)
		}
	}

	userRecommendAreas = make([]*UserRecommendArea, 0)
	for _, vUserRecommendAreaCodes := range userRecommendAreaCodes {
		userRecommendAreas = append(userRecommendAreas, &UserRecommendArea{
			RecommendCode: vUserRecommendAreaCodes,
			Num:           int64(len(strings.Split(vUserRecommendAreaCodes, "D")) - 1),
		})
	}

	if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = uuc.urRepo.CreateUserRecommendArea(ctx, userRecommendAreas)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &v1.CheckAndInsertRecommendAreaReply{}, nil
}

type DownloadRes struct {
	Address   string
	Deposit   uint64
	Buy       uint64
	Withdraw  uint64
	CreatedAt time.Time
}

type DownloadResTwo struct {
	Address       string
	Amount        uint64
	Total         float64
	Reward        float64
	UserCreatedAt time.Time
	CreatedAt     time.Time
}

// internal/biz/user_usecase.go
func (uuc *UserUseCase) BuildDownloadDataExcel(ctx context.Context) (string, string, []byte, error) {
	var (
		recommendUsers []*UserRecommend
		userIds        []int64
		err            error
	)

	recommendUsers, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, "D6D12D13D14D15D621D627D629D630D631D632D633D635D637D638D639D646D658D660D662D669D670D672D673D738D888D898D899")
	if err != nil {
		return "", "", nil, err
	}

	prefix := "D6D12D13D14D15D621D627D629D630D631D632D633D635D637D638D639D646D658D660D662D669D670D672D673D738D888D898D899D900"
	for _, v := range recommendUsers {
		if strings.HasPrefix(v.RecommendCode, prefix) {
			continue
		}
		userIds = append(userIds, v.UserId)
	}

	ethRecords, err := uuc.locationRepo.GetEthUserRecordList(ctx)
	if err != nil {
		return "", "", nil, err
	}

	usersMap, err := uuc.repo.GetUserByUserIdsTwo(ctx, userIds)
	if err != nil {
		return "", "", nil, err
	}

	buyRecords, err := uuc.uiRepo.GetBuyRecordMap(ctx, userIds)
	if err != nil {
		return "", "", nil, err
	}

	withdrawMap, err := uuc.ubRepo.GetWithdrawByUserIdsMap(ctx, userIds)
	if err != nil {
		return "", "", nil, err
	}

	res := make([]*DownloadRes, 0, len(userIds))
	for _, uid := range userIds {
		tmpRes := &DownloadRes{}

		if u, ok := usersMap[uid]; ok {
			tmpRes.Address = u.Address
			tmpRes.CreatedAt = u.CreatedAt
		}

		if list, ok := ethRecords[uid]; ok {
			var sum uint64
			for _, it := range list {
				sum += it.AmountTwo
			}
			tmpRes.Deposit = sum
		}

		if list, ok := buyRecords[uid]; ok {
			var sum uint64
			for _, it := range list {
				sum += uint64(it.Amount)
			}
			tmpRes.Buy = sum
		}

		if list, ok := withdrawMap[uid]; ok {
			var sum uint64
			for _, it := range list {
				sum += uint64(it.AmountNew)
			}
			tmpRes.Withdraw = sum
		}

		res = append(res, tmpRes)
	}

	fileBytes, err := buildDownloadExcel(res)
	if err != nil {
		return "", "", nil, err
	}

	filename := "download_" + time.Now().Format("20060102_150405") + ".xlsx"
	contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	return filename, contentType, fileBytes, nil
}

func buildDownloadExcel(res []*DownloadRes) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName("Sheet1", sheet)

	// 表头
	headers := []string{"地址", "充值", "认购", "提现", "注册时间"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		_ = f.SetCellValue(sheet, cell, h)
	}

	// 数据
	for i, r := range res {
		row := i + 2
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), r.Address)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.Deposit)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.Buy)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.Withdraw)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.CreatedAt)
	}

	// 可选：冻结首行
	_ = f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	// 可选：设置列宽
	_ = f.SetColWidth(sheet, "A", "A", 44)
	_ = f.SetColWidth(sheet, "B", "E", 40)

	// 输出为 bytes
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (uuc *UserUseCase) BuildDownloadDataExcelThree(ctx context.Context) (string, string, []byte, error) {
	var (
		recommendUsers []*UserRecommend
		userIds        []int64
		err            error
	)

	recommendUsers, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, "D6D12D13D14D15D621D627D629D630D631D632D633D635D637D638D639D646D658D660D662D669D670D672D673D738D888D898D899")
	if err != nil {
		return "", "", nil, err
	}

	prefix := "D6D12D13D14D15D621D627D629D630D631D632D633D635D637D638D639D646D658D660D662D669D670D672D673D738D888D898D899D900"
	for _, v := range recommendUsers {
		if strings.HasPrefix(v.RecommendCode, prefix) {
			continue
		}
		userIds = append(userIds, v.UserId)
	}

	usersMap, err := uuc.repo.GetUserByUserIdsTwo(ctx, userIds)
	if err != nil {
		return "", "", nil, err
	}

	buyRecords, err := uuc.uiRepo.GetBuyRecordingMap(ctx, userIds)
	if err != nil {
		return "", "", nil, err
	}

	res := make([]*DownloadResTwo, 0, len(userIds))
	for _, uid := range userIds {

		if list, okT := buyRecords[uid]; okT {

			for _, it := range list {
				tmpRes := &DownloadResTwo{}

				if u, ok := usersMap[uid]; ok {
					tmpRes.Address = u.Address
					tmpRes.UserCreatedAt = u.CreatedAt
				}

				tmpRes.Amount = uint64(it.Amount)
				tmpRes.Reward = it.AmountGet
				tmpRes.Total = it.Amount * 2.5
				tmpRes.CreatedAt = it.CreatedAt
				res = append(res, tmpRes)
			}
		}
	}

	fileBytes, err := buildDownloadExcelThree(res)
	if err != nil {
		return "", "", nil, err
	}

	filename := "download_" + time.Now().Format("20060102_150405") + ".xlsx"
	contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	return filename, contentType, fileBytes, nil
}

func buildDownloadExcelThree(res []*DownloadResTwo) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName("Sheet1", sheet)

	// 表头
	headers := []string{"地址", "用户注册时间", "认购金额", "出局总金额", "已释放金额", "认购时间"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		_ = f.SetCellValue(sheet, cell, h)
	}

	// 数据
	for i, r := range res {
		row := i + 2
		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", row), r.Address)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.UserCreatedAt)
		_ = f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.Amount)
		_ = f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.Total)
		_ = f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.Reward)
		_ = f.SetCellValue(sheet, fmt.Sprintf("F%d", row), r.CreatedAt)
	}

	// 可选：冻结首行
	_ = f.SetPanes(sheet, &excelize.Panes{
		Freeze:      true,
		Split:       false,
		XSplit:      0,
		YSplit:      1,
		TopLeftCell: "A2",
		ActivePane:  "bottomLeft",
	})

	// 可选：设置列宽
	_ = f.SetColWidth(sheet, "A", "A", 44)
	_ = f.SetColWidth(sheet, "B", "F", 40)

	// 输出为 bytes
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (uuc *UserUseCase) VipCheck(ctx context.Context, req *v1.VipCheckRequest) (*v1.VipCheckReply, error) {

	var (
		users           []*UserInfo
		configs         []*Config
		vip5Balance     int64
		vip4Balance     int64
		vip3Balance     int64
		vip2Balance     int64
		vip1Balance     int64
		vip0Balance     int64
		vip5BalanceTeam int64
		vip4BalanceTeam int64
		vip3BalanceTeam int64
		vip2BalanceTeam int64
		vip1BalanceTeam int64
		err             error
	)
	users, err = uuc.repo.GetAllUserInfos(ctx)
	if nil != err {
		return nil, err
	}

	configs, _ = uuc.configRepo.GetConfigByKeys(ctx, "vip_5_balance",
		"vip_4_balance", "vip_3_balance", "vip_2_balance", "vip_1_balance", "vip_0_balance",
		"vip_5_balance_team", "vip_4_balance_team", "vip_3_balance_team", "vip_2_balance_team", "vip_1_balance_team")
	if nil != configs {
		for _, vConfig := range configs {
			if "vip_5_balance" == vConfig.KeyName {
				vip5Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_4_balance" == vConfig.KeyName {
				vip4Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_3_balance" == vConfig.KeyName {
				vip3Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_2_balance" == vConfig.KeyName {
				vip2Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_0_balance" == vConfig.KeyName {
				vip0Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_1_balance" == vConfig.KeyName {
				vip1Balance, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_4_balance_team" == vConfig.KeyName {
				vip4BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_3_balance_team" == vConfig.KeyName {
				vip3BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_2_balance_team" == vConfig.KeyName {
				vip2BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_1_balance_team" == vConfig.KeyName {
				vip1BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			} else if "vip_5_balance_team" == vConfig.KeyName {
				vip5BalanceTeam, _ = strconv.ParseInt(vConfig.Value, 10, 64)
			}
		}
	}

	for _, user := range users {
		if 0 < user.LockVip {
			continue
		}

		var (
			userRecommend  *UserRecommend
			userBalance    *UserBalance
			myCode         string
			teamCsdBalance int64
			myUserBalance  int64
			myVip          int64 = 0
		)

		vip1Count1 := make(map[int64]int64, 0)
		vip2Count1 := make(map[int64]int64, 0)
		vip3Count1 := make(map[int64]int64, 0)
		vip4Count1 := make(map[int64]int64, 0)

		userRecommend, err = uuc.urRepo.GetUserRecommendByUserId(ctx, user.ID)
		if nil != err {
			continue
		}

		userBalance, err = uuc.ubRepo.GetUserBalance(ctx, user.ID)
		if nil != err {
			continue
		}

		// 我的伞下所有用户
		myCode = userRecommend.RecommendCode + "D" + strconv.FormatInt(user.ID, 10)

		var (
			UserInfos             map[int64]*UserInfo
			userRecommends        []*UserRecommend
			userRecommendsUserIds []int64
		)

		userRecommends, err = uuc.urRepo.GetUserRecommendByCode(ctx, myCode)
		if nil == err {
			for _, vUserRecommends := range userRecommends {
				userRecommendsUserIds = append(userRecommendsUserIds, vUserRecommends.UserId)
			}
		}
		if 0 < len(userRecommendsUserIds) {
			UserInfos, err = uuc.uiRepo.GetUserInfoByUserIds(ctx, userRecommendsUserIds...)
		}
		for _, vUserInfos := range UserInfos {
			if 2 == vUserInfos.Vip {
				vip1Count1[vUserInfos.UserId] += 1
			} else if 3 == vUserInfos.Vip {
				vip2Count1[vUserInfos.UserId] += 1
			} else if 4 == vUserInfos.Vip {
				vip3Count1[vUserInfos.UserId] += 1
			} else if 5 == vUserInfos.Vip {
				vip4Count1[vUserInfos.UserId] += 1
			}
		}

		if 0 < len(userRecommends) {
			for _, vUserRecommendsQ := range userRecommends {

				var (
					userRecommends1        []*UserRecommend
					userRecommendsUserIds1 []int64
				)
				myCode1 := vUserRecommendsQ.RecommendCode + "D" + strconv.FormatInt(vUserRecommendsQ.UserId, 10)
				userRecommends1, err = uuc.urRepo.GetUserRecommendLikeCode(ctx, myCode1)
				if nil == err {
					for _, vUserRecommends1 := range userRecommends1 {
						userRecommendsUserIds1 = append(userRecommendsUserIds1, vUserRecommends1.UserId)
					}
				}

				var UserInfos1 map[int64]*UserInfo
				if 0 < len(userRecommendsUserIds1) {
					UserInfos1, err = uuc.uiRepo.GetUserInfoByUserIds(ctx, userRecommendsUserIds1...)
				}

				for _, vUserInfos1 := range UserInfos1 {
					if 2 == vUserInfos1.Vip {
						vip1Count1[vUserRecommendsQ.UserId] += 1
					} else if 3 == vUserInfos1.Vip {
						vip2Count1[vUserRecommendsQ.UserId] += 1
					} else if 4 == vUserInfos1.Vip {
						vip3Count1[vUserRecommendsQ.UserId] += 1
					} else if 5 == vUserInfos1.Vip {
						vip4Count1[vUserRecommendsQ.UserId] += 1
					}
				}
			}
		}

		var (
			vip1Count int64
			vip2Count int64
			vip3Count int64
			vip4Count int64
		)
		for _, vv1 := range vip1Count1 {
			if vv1 > 0 {
				vip1Count++
			}
		}
		for _, vv2 := range vip2Count1 {
			if vv2 > 0 {
				vip2Count++
			}
		}
		for _, vv3 := range vip3Count1 {
			if vv3 > 0 {
				vip3Count++
			}
		}
		for _, vv4 := range vip4Count1 {
			if vv4 > 0 {
				vip4Count++
			}
		}

		teamCsdBalance = user.TeamCsdBalance / 100000
		myUserBalance = userBalance.BalanceUsdt / 100000
		if teamCsdBalance >= vip5BalanceTeam && 2 <= vip4Count && 5 <= user.HistoryRecommend && myUserBalance >= vip5Balance {
			myVip = 6
		} else if teamCsdBalance >= vip4BalanceTeam && 2 <= vip3Count && 5 <= user.HistoryRecommend && myUserBalance >= vip4Balance {
			myVip = 5
		} else if teamCsdBalance >= vip3BalanceTeam && 2 <= vip2Count && 5 <= user.HistoryRecommend && myUserBalance >= vip3Balance {
			myVip = 4
		} else if teamCsdBalance >= vip2BalanceTeam && 2 <= vip1Count && 5 <= user.HistoryRecommend && myUserBalance >= vip2Balance {
			myVip = 3
		} else if teamCsdBalance >= vip1BalanceTeam && 5 <= user.HistoryRecommend && myUserBalance >= vip1Balance {
			myVip = 2
		} else if myUserBalance >= vip0Balance {
			myVip = 1
		}

		if err = uuc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

			// 修改用户推荐人区数据，修改自身区数据
			_, err = uuc.uiRepo.UpdateUserInfoVip(ctx, user.ID, myVip)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return &v1.VipCheckReply{}, nil
}

func (uuc *UserUseCase) CheckAdminUserArea(ctx context.Context, req *v1.CheckAdminUserAreaRequest) (*v1.CheckAdminUserAreaReply, error) {
	return &v1.CheckAdminUserAreaReply{}, nil
}

func (uuc *UserUseCase) CheckAndInsertLocationsRecommendUser(ctx context.Context, req *v1.CheckAndInsertLocationsRecommendUserRequest) (*v1.CheckAndInsertLocationsRecommendUserReply, error) {
	return &v1.CheckAndInsertLocationsRecommendUserReply{}, nil
}

// AdminCreateGoods 处理 HTTP 文件上传请求
func (uuc *UserUseCase) AdminCreateGoods(ctx context.Context, req *v1.AdminCreateGoodsRequest) (*v1.AdminCreateGoodsReply, error) {
	return nil, uuc.repo.UpdateGoods(ctx, req.SendBody.Id, req.SendBody.Status)
}

// AdminCreateGoodsThree 处理 HTTP 文件上传请求
func (uuc *UserUseCase) AdminCreateGoodsThree(ctx context.Context, req *v1.AdminCreateGoodsRequest) (*v1.AdminCreateGoodsReply, error) {
	return nil, uuc.repo.UpdateGoodsThree(ctx, req.SendBody.Id, req.SendBody.Status)
}

// AdminCreateGoodsTwo 处理 HTTP 文件上传请求
func (uuc *UserUseCase) AdminCreateGoodsTwo(ctx context.Context, req *v1.AdminCreateGoodsRequest) (*v1.AdminCreateGoodsReply, error) {
	return nil, uuc.repo.UpdateGoodsTwo(ctx, req.SendBody.Id, req.SendBody.Status)
}

func (uuc *UserUseCase) Upload(ctx transporthttp.Context) (err error) {

	name := ctx.Request().FormValue("name")
	detail := ctx.Request().FormValue("one")
	amount := ctx.Request().FormValue("amount")
	three := ctx.Request().FormValue("three")
	amountInt64, _ := strconv.ParseUint(amount, 10, 64)
	if 0 >= amountInt64 {
		return nil
	}

	if 100 == amountInt64 {
		amountInt64 = 100
	} else if 300 == amountInt64 {
		amountInt64 = 300
	} else if 500 == amountInt64 {
		amountInt64 = 500
	} else if 1000 == amountInt64 {
		amountInt64 = 1000
	} else if 5000 == amountInt64 {
		amountInt64 = 5000
	} else if 10000 == amountInt64 {
		amountInt64 = 10000
	} else if 15000 == amountInt64 {
		amountInt64 = 15000
	} else if 30000 == amountInt64 {
		amountInt64 = 30000
	} else if 50000 == amountInt64 {
		amountInt64 = 50000
	} else if 100000 == amountInt64 {
		amountInt64 = 100000
	} else if 150000 == amountInt64 {
		amountInt64 = 150000
	} else {
		return nil
	}

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()

	picName := time.Now().Format("20060102150405") + ".png"
	// 修改文件名并创建保存图片
	imageFile, err := os.Create("/www/wwwroot/www.ispayplay.com/images/" + picName)
	if err != nil {
		return
	}
	defer imageFile.Close()

	// 将文件内容复制到保存的文件中
	_, err = io.Copy(imageFile, file)
	if err != nil {
		return
	}

	err = uuc.repo.CreateGoods(ctx, detail, name, picName, three, amountInt64)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCase) UploadTwo(ctx transporthttp.Context) (err error) {

	name := ctx.Request().FormValue("name")
	detail := ctx.Request().FormValue("one")
	amount := ctx.Request().FormValue("amount")
	three := ctx.Request().FormValue("three")
	amountInt64, _ := strconv.ParseUint(amount, 10, 64)
	if 0 >= amountInt64 {
		return nil
	}

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()

	picName := time.Now().Format("20060102150405") + ".png"
	// 修改文件名并创建保存图片
	imageFile, err := os.Create("/www/wwwroot/www.ispayplay.com/images/" + picName)
	if err != nil {
		return
	}
	defer imageFile.Close()

	// 将文件内容复制到保存的文件中
	_, err = io.Copy(imageFile, file)
	if err != nil {
		return
	}

	err = uuc.repo.CreateGoodsTwo(ctx, detail, name, picName, three, amountInt64)
	if err != nil {
		return err
	}

	return nil
}

func (uuc *UserUseCase) UploadThree(ctx transporthttp.Context) (err error) {

	name := ctx.Request().FormValue("name")
	detail := ctx.Request().FormValue("one")
	amount := ctx.Request().FormValue("amount")
	three := ctx.Request().FormValue("three")
	amountInt64, _ := strconv.ParseUint(amount, 10, 64)
	if 0 >= amountInt64 {
		return nil
	}

	file, _, err := ctx.Request().FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()

	picName := time.Now().Format("20060102150405") + ".png"
	// 修改文件名并创建保存图片
	imageFile, err := os.Create("/www/wwwroot/www.ispayplay.com/images/" + picName)
	if err != nil {
		return
	}
	defer imageFile.Close()

	// 将文件内容复制到保存的文件中
	_, err = io.Copy(imageFile, file)
	if err != nil {
		return
	}

	err = uuc.repo.CreateGoodsThree(ctx, detail, name, picName, three, amountInt64)
	if err != nil {
		return err
	}

	return nil
}
